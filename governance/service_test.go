package governance_test

import (
	"context"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/governance"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/orders/mocks"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

type testSvc struct {
	*governance.Svc
	ctrl  *gomock.Controller
	ctx   context.Context
	cfunc context.CancelFunc

	time *mocks.MockTimeService
}

func newTestService(t *testing.T) *testSvc {
	ctrl := gomock.NewController(t)
	time := mocks.NewMockTimeService(ctrl)

	ctx, cfunc := context.WithCancel(context.Background())

	result := &testSvc{
		ctrl:  ctrl,
		ctx:   ctx,
		cfunc: cfunc,
		time:  time,
	}
	result.Svc = governance.NewService(logging.NewTestLogger(), governance.NewDefaultConfig(), time)
	assert.NotNil(t, result.Svc)
	return result
}

func TestGovernanceService(t *testing.T) {
	svc := newTestService(t)

	cfg := svc.Config
	cfg.Level.Level = logging.DebugLevel
	svc.ReloadConf(cfg)
	assert.Equal(t, svc.Config.Level.Level, logging.DebugLevel)

	cfg.Level.Level = logging.InfoLevel
	svc.ReloadConf(cfg)
	assert.Equal(t, svc.Config.Level.Level, logging.InfoLevel)
}

func TestPrepareProposal(t *testing.T) {
	svc := newTestService(t)

	updateNetwork := types.Proposal_Terms_UpdateNetwork{
		Changes: &types.NetworkConfiguration{
			MinCloseInDays: 100,
			MaxCloseInDays: 1000,
		},
	}
	terms := types.Proposal_Terms{
		Parameters: &types.Proposal_Terms_Parameters{
			CloseInDays:           30,
			EnactInDays:           31,
			MinParticipationStake: 50,
		},
		Change: &types.Proposal_Terms_UpdateNetwork_{
			UpdateNetwork: &updateNetwork,
		},
	}

	rightNow := time.Now()
	svc.time.EXPECT().GetTimeNow().Times(1).Return(rightNow, nil)

	testAuthor := "test-author"
	proposal, err := svc.PrepareProposal(svc.ctx, testAuthor, "", &terms)

	assert.NoError(t, err)
	assert.NotNil(t, proposal)
	assert.NotEmpty(t, proposal.Reference, "reference expected to be auto-generated if empty")
	assert.EqualValues(t, testAuthor, proposal.PartyID)
	assert.EqualValues(t, types.Proposal_OPEN, proposal.State)
	assert.EqualValues(t, terms, *proposal.Terms)
}

func TestPrepareEmptyProposal(t *testing.T) {
	svc := newTestService(t)

	updateNetwork := types.Proposal_Terms_UpdateNetwork{
		Changes: &types.NetworkConfiguration{},
	}
	terms := types.Proposal_Terms{
		Parameters: &types.Proposal_Terms_Parameters{},
		Change: &types.Proposal_Terms_UpdateNetwork_{
			UpdateNetwork: &updateNetwork,
		},
	}

	svc.time.EXPECT().GetTimeNow().MaxTimes(0)

	proposal, err := svc.PrepareProposal(svc.ctx, "", "", &terms)

	assert.Error(t, err)
	assert.Nil(t, proposal)
	assert.Contains(t, err.Error(), "proposal validation failed")
}

func TestValidateTerms(t *testing.T) {
	svc := newTestService(t)

	updateNetwork := types.Proposal_Terms_UpdateNetwork{
		Changes: &types.NetworkConfiguration{},
	}

	err := svc.ValidateTerms(&types.Proposal_Terms{
		Parameters: &types.Proposal_Terms_Parameters{},
		Change: &types.Proposal_Terms_UpdateNetwork_{
			UpdateNetwork: &updateNetwork,
		},
	})
	assert.Error(t, err)
}
