package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	"github.com/golang/mock/gomock"
	testutilctx "github.com/obada-foundation/fullcore/testutil/context"
	testutilkeepers "github.com/obada-foundation/fullcore/testutil/keeper/nft"
	testutilcodec "github.com/obada-foundation/fullcore/testutil/module/codec"
	simtestutil "github.com/obada-foundation/fullcore/testutil/sims"
	module "github.com/obada-foundation/fullcore/x/obit"
	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

var class = types.Class{
	Id:     "OBT",
	Name:   "test name",
	Symbol: "TST",
	Uri:    "https://obada.io",
}

// KeeperTestSuite is the keeper test suite
type KeeperTestSuite struct {
	suite.Suite
	addrs []sdk.AccAddress
	ctx   sdk.Context

	msgServer  types.MsgServer
	obitKeeper keeper.Keeper
	nftKeeper  nftkeeper.Keeper

	encCfg testutilcodec.TestEncodingConfig
}

// SetupTest creates a test suite to test the obit module
func (s *KeeperTestSuite) SetupTest() {
	s.addrs = simtestutil.CreateIncrementalAccounts(3)
	key := sdk.NewKVStoreKey(types.StoreKey)
	memKey := sdk.NewKVStoreKey(types.MemStoreKey)

	s.encCfg = testutilcodec.MakeTestEncodingConfig(module.AppModuleBasic{})
	nftKey := sdk.NewKVStoreKey(nft.StoreKey)

	ctx := testutilctx.DefaultContextWithDB(
		s.T(),
		nftKey,
		sdk.NewTransientStoreKey("transient_test"),
	).Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	accountKeeper := testutilkeepers.NewMockAccountKeeper(ctrl)
	bankKeeper := testutilkeepers.NewMockBankKeeper(ctrl)

	accountKeeper.EXPECT().GetModuleAddress("nft").Return(s.addrs[0]).AnyTimes()

	nftKeeper := nftkeeper.NewKeeper(
		nftKey,
		s.encCfg.Codec,
		accountKeeper,
		bankKeeper,
	)

	obitKeeper := keeper.NewKeeper(s.encCfg.Codec, key, memKey, nftKeeper)

	s.obitKeeper = *obitKeeper
	s.nftKeeper = nftKeeper
	s.msgServer = keeper.NewMsgServerImpl(*obitKeeper)
	s.ctx = ctx
}

// TestKeeperTestSuite runs the keeper test suite
func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

// TestSaveClass tests saving a class
func (s *KeeperTestSuite) TestSaveClass() {
	s.Assert().False(s.nftKeeper.HasClass(s.ctx, class.Id))

	err := s.obitKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	s.Assert().True(s.nftKeeper.HasClass(s.ctx, class.Id))
}
