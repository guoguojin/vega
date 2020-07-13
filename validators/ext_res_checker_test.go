package validators_test

import (
	"testing"
	"time"

	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/proto"
	"code.vegaprotocol.io/vega/validators"
	"code.vegaprotocol.io/vega/validators/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testExtResChecker struct {
	*validators.ExtResChecker
	ctrl      *gomock.Controller
	top       *mocks.MockValidatorTopology
	cmd       *mocks.MockCommander
	tsvc      *mocks.MockTimeService
	startTime time.Time
}

func getTestExtResChecker(t *testing.T) *testExtResChecker {
	ctrl := gomock.NewController(t)
	top := mocks.NewMockValidatorTopology(ctrl)
	cmd := mocks.NewMockCommander(ctrl)
	tsvc := mocks.NewMockTimeService(ctrl)

	now := time.Now()
	tsvc.EXPECT().GetTimeNow().Times(1).Return(now, nil)
	tsvc.EXPECT().NotifyOnTick(gomock.Any()).Times(1)
	nv := validators.NewExtResChecker(
		logging.NewTestLogger(), top, cmd, tsvc)
	assert.NotNil(t, nv)

	return &testExtResChecker{
		ExtResChecker: nv,
		ctrl:          ctrl,
		top:           top,
		cmd:           cmd,
		tsvc:          tsvc,
		startTime:     now,
	}
}

func TestExtResCheck(t *testing.T) {
	t.Run("start - error duplicate", testStartErrorDuplicate)
	t.Run("start - error check failed", testStartErrorCheckFailed)
	t.Run("start - OK", testStartOK)
	t.Run("add node vote - error invalid id", testNodeVoteInvalidProposalReference)
	t.Run("add node vote - error note a validator", testNodeVoteNotAValidator)
	t.Run("add node vote - error duplicate vote", testNodeVoteDuplicateVote)
	t.Run("add node vote - OK", testNodeVoteOK)
	t.Run("on chain time update validated asset", testOnChainTimeUpdate)
}

func testStartErrorDuplicate(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)
	err = erc.StartCheck(res, cb, checkUntil)
	assert.EqualError(t, err, validators.ErrResourceDuplicate.Error())
}

func testStartErrorCheckFailed(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(1 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.EqualError(t, err, validators.ErrCheckUntilInvalid.Error())
}

func testStartOK(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)
}

func testNodeVoteInvalidProposalReference(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)

	err = erc.AddNodeCheck(&proto.NodeVote{Reference: "bad-id"})
	assert.EqualError(t, err, validators.ErrInvalidResourceIDForNodeVote.Error())
}

func testNodeVoteNotAValidator(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)

	erc.top.EXPECT().Exists(gomock.Any()).Times(1).Return(false)
	err = erc.AddNodeCheck(&proto.NodeVote{Reference: res.id})
	assert.EqualError(t, err, validators.ErrVoteFromNonValidator.Error())
}

func testNodeVoteOK(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)

	erc.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
	err = erc.AddNodeCheck(&proto.NodeVote{Reference: res.id})
	assert.NoError(t, err)
}

func testNodeVoteDuplicateVote(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)

	// first vote, all good
	erc.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
	err = erc.AddNodeCheck(&proto.NodeVote{Reference: res.id, PubKey: []byte("somepubkey")})
	assert.NoError(t, err)

	// second vote, bad
	erc.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
	err = erc.AddNodeCheck(&proto.NodeVote{Reference: res.id, PubKey: []byte("somepubkey")})
	assert.EqualError(t, err, validators.ErrDuplicateVoteFromNode.Error())
}

func testOnChainTimeUpdate(t *testing.T) {
	erc := getTestExtResChecker(t)
	defer erc.ctrl.Finish()
	defer erc.Stop()

	selfPubKey := []byte("selfPubKey")

	erc.top.EXPECT().Len().AnyTimes().Return(2)
	erc.top.EXPECT().IsValidator().AnyTimes().Return(true)
	erc.top.EXPECT().SelfVegaPubKey().AnyTimes().Return(selfPubKey)

	ch := make(chan struct{}, 1)
	res := testRes{"resource-id-1", func() error {
		return nil
	}}
	checkUntil := erc.startTime.Add(700 * time.Second)
	cb := func(interface{}, bool) {
		// unblock chanel listen to finish test
		ch <- struct{}{}
	}

	err := erc.StartCheck(res, cb, checkUntil)
	assert.NoError(t, err)

	// first on chain time update, we send our own vote
	erc.cmd.EXPECT().Command(gomock.Any(), gomock.Any()).Times(1).Return(nil)
	newNow := erc.startTime.Add(1 * time.Second)
	erc.OnTick(newNow)

	// then we propagate our own vote
	erc.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
	err = erc.AddNodeCheck(&proto.NodeVote{Reference: res.id, PubKey: selfPubKey})
	assert.NoError(t, err)

	// second vote from another validator
	erc.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
	err = erc.AddNodeCheck(&proto.NodeVote{Reference: res.id, PubKey: []byte("somepubkey")})
	assert.NoError(t, err)

	// call onTick again to get the callback called
	newNow = newNow.Add(1 * time.Second)
	erc.OnTick(newNow)

	// block to wait for the result
	<-ch
}

type testRes struct {
	id    string
	check func() error
}

func (t testRes) GetID() string { return t.id }
func (t testRes) Check() error  { return t.check() }

// func testOnChainTimeUpdate(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	// first submit a proposal
// 	builtinAsset := &types.BuiltinAsset{
// 		Name:   "USDC",
// 		Symbol: "USDC",
// 	}

// 	now := time.Now()
// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Reference: "REF",
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_BuiltinAsset{
// 							BuiltinAsset: builtinAsset,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(
// 		gomock.Any(), gomock.Any()).
// 		Times(1).Return("ASSETID", nil)

// 	ch := make(chan struct{}, 1)
// 	asset := builtin.New("ASSETID", builtinAsset)
// 	nv.assets.EXPECT().Get(gomock.Any()).Times(1).DoAndReturn(func(string) (assets.Asset, error) {
// 		ch <- struct{}{}
// 		return asset, nil
// 	})

// 	err := nv.Start(p)
// 	<-ch
// 	assert.NoError(t, err)

// 	// here we gonna wait just a little because hte validation is done asynchronously, and while we use a channel previously to ensure that the function to get the asset we not sure we ran through the validation already
// 	time.Sleep(50 * time.Millisecond)

// 	nv.top.EXPECT().Len().AnyTimes().Return(1)
// 	nv.top.EXPECT().SelfVegaPubKey().AnyTimes().Return([]byte("okkey"))
// 	nv.cmd.EXPECT().Command(gomock.Any(), gomock.Any()).Times(1).Return(nil)

// 	// no we call time update once.
// 	// this will call the commander to send the NodeVote
// 	// and return 0 validatedProposals
// 	validatedProposals, _ := nv.OnChainTimeUpdate(now.Add(1 * time.Second))
// 	assert.Len(t, validatedProposals, 0)

// 	// no we submit the vote

// 	v := &types.NodeVote{
// 		Reference: "REF",
// 		PubKey:    []byte("avalidator"),
// 	}

// 	nv.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
// 	err = nv.AddNodeVote(v)
// 	assert.NoError(t, err)

// 	//we call time update once more now that the proposal should have all node votes
// 	validatedProposals, _ = nv.OnChainTimeUpdate(now.Add(1 * time.Second))
// 	assert.Len(t, validatedProposals, 1)
// }

// func testNodeValidationRequiredTrue(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	p := &types.Proposal{
// 		Terms: &types.ProposalTerms{
// 			Change: &types.ProposalTerms_NewAsset{},
// 		},
// 	}

// 	assert.True(t, nv.IsNodeValidationRequired(p))
// }

// func testNodeValidationRequiredFalse(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	p := &types.Proposal{
// 		Terms: &types.ProposalTerms{
// 			Change: &types.ProposalTerms_NewMarket{},
// 		},
// 	}

// 	assert.False(t, nv.IsNodeValidationRequired(p))
// }

// func testStartErrorNoNodeValidationRequired(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	p := &types.Proposal{
// 		Terms: &types.ProposalTerms{
// 			Change: &types.ProposalTerms_NewMarket{},
// 		},
// 	}

// 	err := nv.Start(p)
// 	assert.EqualError(t, err, governance.ErrNoNodeValidationRequired.Error())
// }

// func testNodeVoteInvalidProposalReference(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	v := &types.NodeVote{
// 		Reference: "nope",
// 	}

// 	err := nv.AddNodeVote(v)
// 	assert.EqualError(t, err, governance.ErrInvalidProposalReferenceForNodeVote.Error())
// }

// func testNodeVoteNotAValidator(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	// first submit a proposal
// 	builtinAsset := &types.BuiltinAsset{
// 		Name:   "USDC",
// 		Symbol: "USDC",
// 	}

// 	now := time.Now()
// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Reference: "REF",
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_BuiltinAsset{
// 							BuiltinAsset: builtinAsset,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(
// 		gomock.Any(), gomock.Any()).
// 		Times(1).Return("ASSETID", nil)

// 	ch := make(chan struct{}, 1)
// 	asset := builtin.New("ASSETID", builtinAsset)
// 	nv.assets.EXPECT().Get(gomock.Any()).Times(1).DoAndReturn(func(string) (assets.Asset, error) {
// 		ch <- struct{}{}
// 		return asset, nil
// 	})

// 	err := nv.Start(p)
// 	<-ch
// 	assert.NoError(t, err)

// 	// no we submit the vote

// 	v := &types.NodeVote{
// 		Reference: "REF",
// 		PubKey:    []byte("notavalidator"),
// 	}

// 	nv.top.EXPECT().Exists(gomock.Any()).Times(1).Return(false)
// 	err = nv.AddNodeVote(v)
// 	assert.EqualError(t, err, governance.ErrNodeIsNotAValidator.Error())
// }

// func testNodeVoteDuplicateVote(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	// first submit a proposal
// 	builtinAsset := &types.BuiltinAsset{
// 		Name:   "USDC",
// 		Symbol: "USDC",
// 	}

// 	now := time.Now()
// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Reference: "REF",
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_BuiltinAsset{
// 							BuiltinAsset: builtinAsset,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(
// 		gomock.Any(), gomock.Any()).
// 		Times(1).Return("ASSETID", nil)

// 	ch := make(chan struct{}, 1)
// 	asset := builtin.New("ASSETID", builtinAsset)
// 	nv.assets.EXPECT().Get(gomock.Any()).Times(1).DoAndReturn(func(string) (assets.Asset, error) {
// 		ch <- struct{}{}
// 		return asset, nil
// 	})

// 	err := nv.Start(p)
// 	<-ch
// 	assert.NoError(t, err)

// 	// no we submit the vote

// 	v := &types.NodeVote{
// 		Reference: "REF",
// 		PubKey:    []byte("notavalidator"),
// 	}

// 	nv.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
// 	err = nv.AddNodeVote(v)
// 	assert.NoError(t, err)

// 	nv.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
// 	err = nv.AddNodeVote(v)
// 	assert.EqualError(t, err, governance.ErrDuplicateVoteFromNode.Error())
// }

// func testNodeVoteOK(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	// first submit a proposal
// 	builtinAsset := &types.BuiltinAsset{
// 		Name:   "USDC",
// 		Symbol: "USDC",
// 	}

// 	now := time.Now()
// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Reference: "REF",
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_BuiltinAsset{
// 							BuiltinAsset: builtinAsset,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(
// 		gomock.Any(), gomock.Any()).
// 		Times(1).Return("ASSETID", nil)

// 	ch := make(chan struct{}, 1)
// 	asset := builtin.New("ASSETID", builtinAsset)
// 	nv.assets.EXPECT().Get(gomock.Any()).Times(1).DoAndReturn(func(string) (assets.Asset, error) {
// 		ch <- struct{}{}
// 		return asset, nil
// 	})

// 	err := nv.Start(p)
// 	<-ch
// 	assert.NoError(t, err)

// 	// no we submit the vote

// 	v := &types.NodeVote{
// 		Reference: "REF",
// 		PubKey:    []byte("notavalidator"),
// 	}

// 	nv.top.EXPECT().Exists(gomock.Any()).Times(1).Return(true)
// 	err = nv.AddNodeVote(v)
// 	assert.NoError(t, err)
// }

// func testStartErrorCheckProposalFailed(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    1,
// 			ValidationTimestamp: 2,
// 			Change:              &types.ProposalTerms_NewAsset{},
// 		},
// 	}

// 	err := nv.Start(p)
// 	assert.EqualError(t, err, governance.ErrProposalValidationTimestampInvalid.Error())

// 	// now both are under required duration
// 	p.Terms.ClosingTimestamp = 3
// 	err = nv.Start(p)
// 	assert.EqualError(t, err, governance.ErrProposalValidationTimestampInvalid.Error())

// }

// func testStartOK(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	now := time.Now()

// 	builtinAsset := &types.BuiltinAsset{
// 		Name:   "USDC",
// 		Symbol: "USDC",
// 	}

// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_BuiltinAsset{
// 							BuiltinAsset: builtinAsset,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(
// 		gomock.Any(), gomock.Any()).
// 		Times(1).Return("ASSETID", nil)

// 	ch := make(chan struct{}, 1)
// 	asset := builtin.New("ASSETID", builtinAsset)
// 	nv.assets.EXPECT().Get(gomock.Any()).Times(1).DoAndReturn(func(string) (assets.Asset, error) {
// 		ch <- struct{}{}
// 		return asset, nil
// 	})

// 	err := nv.Start(p)
// 	<-ch
// 	assert.NoError(t, err)

// }

// func testStartErrorDuplicate(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	now := time.Now()

// 	builtinAsset := &types.BuiltinAsset{
// 		Name:   "USDC",
// 		Symbol: "USDC",
// 	}

// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Reference: "ref",
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_BuiltinAsset{
// 							BuiltinAsset: builtinAsset,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(
// 		gomock.Any(), gomock.Any()).
// 		Times(1).Return("ASSETID", nil)

// 	ch := make(chan struct{}, 1)
// 	asset := builtin.New("ASSETID", builtinAsset)
// 	nv.assets.EXPECT().Get(gomock.Any()).Times(1).DoAndReturn(func(string) (assets.Asset, error) {
// 		ch <- struct{}{}
// 		return asset, nil
// 	})

// 	err := nv.Start(p)
// 	<-ch
// 	assert.NoError(t, err)

// 	// first was fine, now let's try to submit it again
// 	err = nv.Start(p)
// 	assert.EqualError(t, err, governance.ErrProposalReferenceDuplicate.Error())
// }

// func testStartErrorUnableToInstanciateAsset(t *testing.T) {
// 	nv := getTestNodeValidation(t)
// 	defer nv.ctrl.Finish()

// 	now := time.Now()

// 	// first closing time < validation time
// 	p := &types.Proposal{
// 		Terms: &types.ProposalTerms{
// 			ClosingTimestamp:    now.Add(24 * time.Hour).Unix(),
// 			ValidationTimestamp: now.Add(700 * time.Minute).Unix(),
// 			Change: &types.ProposalTerms_NewAsset{
// 				NewAsset: &types.NewAsset{
// 					Changes: &types.AssetSource{
// 						Source: &types.AssetSource_Erc20{
// 							Erc20: &types.ERC20{
// 								ContractAddress: "0xOK",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	nv.assets.EXPECT().NewAsset(gomock.Any(), gomock.Any()).Times(1).Return("", errors.New("unable to instanciate"))

// 	err := nv.Start(p)
// 	assert.EqualError(t, err, "unable to instanciate")
// }
