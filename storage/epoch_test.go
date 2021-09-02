package storage_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"code.vegaprotocol.io/data-node/logging"
	"code.vegaprotocol.io/data-node/storage"
	pb "code.vegaprotocol.io/protos/vega"
	"github.com/stretchr/testify/assert"
)

func TestEpochs(t *testing.T) {
	a := assert.New(t)

	log := logging.NewTestLogger()
	c := storage.NewDefaultConfig("")

	nodeStore := storage.NewNode(log, c)
	epochStore := storage.NewEpoch(log, nodeStore, c)

	epoch, err := epochStore.GetEpoch()
	a.EqualError(err, "no epoch present")
	a.Nil(epoch)

	epoch, err = epochStore.GetEpochByID("epoch_id")
	a.EqualError(err, "epoch epoch_id not found")
	a.Nil(epoch)

	startTime := time.Date(2020, time.December, 25, 12, 0, 0, 0, time.UTC)
	endTime := startTime.Add(24 * time.Hour)

	epochStore.AddEpoch(1, startTime.Unix(), endTime.Unix())

	epoch, err = epochStore.GetEpochByID("1")
	a.NoError(err)
	assertEpoch(a, epoch, []*pb.Delegation{}, []*pb.Node{}, 1, startTime.Unix(), endTime.Unix())

	epoch, err = epochStore.GetEpoch()
	a.NoError(err)
	assertEpoch(a, epoch, []*pb.Delegation{}, []*pb.Node{}, 1, startTime.Unix(), endTime.Unix())

	delegations := []*pb.Delegation{
		{
			EpochSeq: "1",
			Party:    "party_1",
			NodeId:   "node_1",
			Amount:   "10",
		},
		{
			EpochSeq: "1",
			Party:    "party_2",
			NodeId:   "node_2",
			Amount:   "5",
		},
	}

	// Add delegations to existing epoch
	epochStore.AddDelegation(*delegations[0])
	epochStore.AddDelegation(*delegations[1])

	epoch, err = epochStore.GetEpoch()
	a.NoError(err)
	assertEpoch(a, epoch, delegations, []*pb.Node{}, 1, startTime.Unix(), endTime.Unix())

	// Add delegations to epoch that hasn't arrived yet
	delegations[0].EpochSeq = "2"
	delegations[1].EpochSeq = "2"

	epochStore.AddDelegation(*delegations[0])
	epochStore.AddDelegation(*delegations[1])

	epoch, err = epochStore.GetEpochByID("2")
	a.NoError(err)
	assertEpoch(a, epoch, delegations, []*pb.Node{}, 2, 0, 0)

	// Add epoch that already holds delegations - this will update the epoch
	startTime = startTime.Add(24 * time.Hour)
	endTime = endTime.Add(24 * time.Hour)
	epochStore.AddEpoch(2, startTime.Unix(), endTime.Unix())

	epoch, err = epochStore.GetEpochByID("2")
	a.NoError(err)
	assertEpoch(a, epoch, delegations, []*pb.Node{}, 2, startTime.Unix(), endTime.Unix())
	epoch, err = epochStore.GetEpoch()
	a.NoError(err)
	assertEpoch(a, epoch, delegations, []*pb.Node{}, 2, startTime.Unix(), endTime.Unix())

	uptime := epochStore.GetTotalNodesUptime()
	a.Equal(48*time.Hour, uptime)

	var nodes []*pb.Node
	for i := 0; i < 2; i++ {
		nodes = append(nodes, &pb.Node{
			Id:                fmt.Sprintf("%d", i),
			PubKey:            fmt.Sprintf("pub_key_%d", i),
			InfoUrl:           fmt.Sprintf("node-%d.xyz.vega/info", i),
			Location:          "GB",
			Status:            pb.NodeStatus_NODE_STATUS_VALIDATOR,
			StakedByOperator:  "0",
			StakedByDelegates: "0",
			StakedTotal:       "0",
			Delagations:       nil,
		})
	}

	// Test epoch returns nodes
	nodeStore.AddNode(*nodes[0])
	nodeStore.AddNode(*nodes[1])

	startTime = startTime.Add(24 * time.Hour)
	endTime = endTime.Add(24 * time.Hour)

	epochStore.AddEpoch(3, startTime.Unix(), endTime.Unix())

	delegations[0].EpochSeq = "3"
	delegations[1].EpochSeq = "3"
	epochStore.AddDelegation(*delegations[0])
	epochStore.AddDelegation(*delegations[1])

	epoch, err = epochStore.GetEpoch()
	a.NoError(err)
	assertEpoch(a, epoch, delegations, nodes, 3, startTime.Unix(), endTime.Unix())

	a.Equal(72*time.Hour, epochStore.GetTotalNodesUptime())
}

func assertEpoch(
	a *assert.Assertions,
	epoch *pb.Epoch,
	delegations []*pb.Delegation,
	nodes []*pb.Node,
	seq uint64,
	startTime, endTime int64,
) {
	a.Equal(epoch.Seq, seq)
	a.Equal(epoch.Timestamps.StartTime, startTime)
	a.Equal(epoch.Timestamps.EndTime, endTime)

	a.Equal(len(delegations), len(epoch.Delegations))

	sort.Sort(ByXY(delegations))
	sort.Sort(ByXY(epoch.Delegations))

	for i := range delegations {
		a.Equal(delegations[i], epoch.Delegations[i])
	}

	a.Equal(len(nodes), len(epoch.Validators))

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Id < nodes[j].Id
	})
	sort.Slice(epoch.Validators, func(i, j int) bool {
		return epoch.Validators[i].Id < epoch.Validators[j].Id
	})

	for i := range nodes {
		a.Equal(nodes[i], epoch.Validators[i])
	}
}
