package types

import (
	"encoding/hex"
	"errors"

	"code.vegaprotocol.io/vega/libs/crypto"

	snapshot "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"github.com/cosmos/iavl"
	"github.com/golang/protobuf/proto"
	tmtypes "github.com/tendermint/tendermint/abci/types"
)

type SnapshotNamespace string

const (
	undefinedSnapshot  SnapshotNamespace = ""
	AppSnapshot        SnapshotNamespace = "app"
	AssetsSnapshot     SnapshotNamespace = "assets"
	BankingSnapshot    SnapshotNamespace = "banking"
	CheckpointSnapshot SnapshotNamespace = "checkpoint"
	CollateralSnapshot SnapshotNamespace = "collateral"
	NetParamsSnapshot  SnapshotNamespace = "netparams"
	DelegationSnapshot SnapshotNamespace = "delegation"
	GovernanceSnapshot SnapshotNamespace = "governance"
	PositionsSnapshot  SnapshotNamespace = "positions"
	MatchingSnapshot   SnapshotNamespace = "matching"
	ExecutionSnapshot  SnapshotNamespace = "execution"
	EpochSnapshot      SnapshotNamespace = "epoch"
	StakingSnapshot    SnapshotNamespace = "staking"

	MaxChunkSize   = 16 * 1000 * 1000 // technically 16 * 1024 * 1024, but you know
	IdealChunkSize = 10 * 1000 * 1000 // aim for 10MB
)

var (
	nsMap = map[string]SnapshotNamespace{
		"collateral": CollateralSnapshot,
		"assets":     AssetsSnapshot,
		"banking":    BankingSnapshot,
		"checkpoint": CheckpointSnapshot,
		"app":        AppSnapshot,
		"netparams":  NetParamsSnapshot,
		"delegation": DelegationSnapshot,
		"governance": GovernanceSnapshot,
		"positions":  PositionsSnapshot,
		"matching":   MatchingSnapshot,
		"execution":  ExecutionSnapshot,
		"epoch":      EpochSnapshot,
		"staking":    StakingSnapshot,
	}

	ErrUnknownSnapshotNamespace  = errors.New("unknown snapshot namespace")
	ErrNoPrefixFound             = errors.New("no prefix in chunk keys")
	ErrInconsistentNamespaceKeys = errors.New("chunk contains several namespace keys")
	ErrChunkHashMismatch         = errors.New("loaded chunk hash does not match metadata")
	ErrChunkOutOfRange           = errors.New("chunk number out of range")
)

type SnapshotFormat = snapshot.Format

const (
	SnapshotFormatUnspecified     = snapshot.Format_FORMAT_UNSPECIFIED
	SnapshotFormatProto           = snapshot.Format_FORMAT_PROTO
	SnapshotFormatProtoCompressed = snapshot.Format_FORMAT_PROTO_COMPRESSED
	SnapshotFormatJSON            = snapshot.Format_FORMAT_JSON
)

func SnapshotFromTM(tms *tmtypes.Snapshot) (*Snapshot, error) {
	snap := Snapshot{
		Height:     tms.Height,
		Format:     SnapshotFormat(tms.Format),
		Chunks:     tms.Chunks,
		Hash:       tms.Hash,
		Metadata:   tms.Metadata,
		ByteChunks: make([][]byte, int(tms.Chunks)), // have the chunk slice ready for loading
	}
	meta := &snapshot.Metadata{}
	if err := proto.Unmarshal(tms.Metadata, meta); err != nil {
		return nil, err
	}
	md, err := MetadataFromProto(meta)
	if err != nil {
		return nil, err
	}
	snap.Meta = md
	return &snap, nil
}

func SnapshotFromIAVL(tree *iavl.ImmutableTree, keys []string) (*Snapshot, error) {
	snap := Snapshot{
		Hash: tree.Hash(),
		Meta: &Metadata{
			Version:     tree.Version(),
			NodeHashes:  make([]*NodeHash, 0, len(keys)),
			ChunkHashes: make([]string, 0, len(keys)), // this is probably premature
		},
		Nodes: make([]*Payload, 0, len(keys)), // each node as a payload
	}
	for _, k := range keys {
		_, val := tree.Get([]byte(k))
		pl := &snapshot.Payload{}
		if err := proto.Unmarshal(val, pl); err != nil {
			return nil, err
		}
		payload := PayloadFromProto(pl)
		payload.raw = val
		hash := hex.EncodeToString(crypto.Hash(val))
		nh := &NodeHash{
			FullKey:   k,
			Namespace: payload.Namespace(),
			Key:       payload.Key(),
			Hash:      hash,
		}
		snap.Meta.NodeHashes = append(snap.Meta.NodeHashes, nh)
		snap.Nodes = append(snap.Nodes, PayloadFromProto(pl))
	}
	// divide into chunks, and set the meta...
	snap.nodesToChunks()
	return &snap, nil
}

func (s *Snapshot) nodesToChunks() {
	all := &Chunk{
		Data: s.Nodes[:],
		Nr:   1,
		Of:   1,
	}
	b, _ := proto.Marshal(all.IntoProto())
	if len(b) < MaxChunkSize {
		s.DataChunks = []*Chunk{
			all,
		}
		s.hashChunks()
		return
	}
	parts := len(b) / IdealChunkSize
	if t := parts * IdealChunkSize; t != len(b) {
		parts++
	}
	s.ByteChunks = make([][]byte, 0, parts)
	step := len(b) / parts
	for i := 0; i < len(b); i += step {
		end := i + step
		if end > len(b) {
			end = len(b)
		}
		s.ByteChunks = append(s.ByteChunks, b[i:end])
	}
	s.hashByteChunks()
}

func (s *Snapshot) hashByteChunks() {
	s.Meta.ChunkHashes = make([]string, 0, len(s.ByteChunks))
	for _, b := range s.ByteChunks {
		s.Meta.ChunkHashes = append(s.Meta.ChunkHashes, hex.EncodeToString(crypto.Hash(b)))
		s.Chunks++
	}
}

func (s *Snapshot) hashChunks() {
	s.Meta.ChunkHashes = make([]string, 0, len(s.DataChunks))
	s.ByteChunks = make([][]byte, 0, len(s.DataChunks))
	for _, c := range s.DataChunks {
		pc := c.IntoProto()
		b, _ := proto.Marshal(pc)
		s.Meta.ChunkHashes = append(s.Meta.ChunkHashes, hex.EncodeToString(crypto.Hash(b)))
		s.ByteChunks = append(s.ByteChunks, b)
		s.Chunks++
	}
}

func (s *Snapshot) LoadChunk(nr uint32, data []byte) error {
	if nr > s.Chunks {
		return ErrChunkOutOfRange
	}
	if len(s.ByteChunks) == 0 {
		s.ByteChunks = make([][]byte, int(s.Chunks))
	}
	i := int(nr)
	if len(s.Meta.ChunkHashes) <= i {
		return ErrChunkOutOfRange
	}
	hash := hex.EncodeToString(crypto.Hash(data))
	if s.Meta.ChunkHashes[i] != hash {
		return ErrChunkHashMismatch
	}
	s.ByteChunks[i] = data
	s.byteLen += len(data)
	s.ChunksSeen += 1
	if s.Chunks == s.ChunksSeen {
		return s.unmarshalChunks()
	}
	return nil
}

func (s *Snapshot) unmarshalChunks() error {
	data := make([]byte, 0, s.byteLen)
	for _, b := range s.ByteChunks {
		data = append(data, b...)
	}
	sChunk := &snapshot.Chunk{}
	if err := proto.Unmarshal(data, sChunk); err != nil {
		return err
	}
	s.DataChunks = []*Chunk{
		ChunkFromProto(sChunk),
	}
	return nil
}

func (s Snapshot) Ready() bool {
	return s.ChunksSeen == s.Chunks
}

func namespaceFromString(s string) (SnapshotNamespace, error) {
	ns, ok := nsMap[s]
	if !ok {
		return undefinedSnapshot, ErrUnknownSnapshotNamespace
	}
	return ns, nil
}

func (n SnapshotNamespace) String() string {
	return string(n)
}
