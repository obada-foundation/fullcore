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

type KeeperTestSuite struct {
	suite.Suite
	addrs []sdk.AccAddress
	ctx   sdk.Context

	msgServer  types.MsgServer
	obitKeeper keeper.Keeper
	nftKeeper  nftkeeper.Keeper

	encCfg testutilcodec.TestEncodingConfig
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.addrs = simtestutil.CreateIncrementalAccounts(3)
	key := sdk.NewKVStoreKey(types.StoreKey)
	memKey := sdk.NewKVStoreKey(types.MemStoreKey)

	suite.encCfg = testutilcodec.MakeTestEncodingConfig(module.AppModuleBasic{})
	nftKey := sdk.NewKVStoreKey(nft.StoreKey)

	ctx := testutilctx.DefaultContextWithDB(
		suite.T(),
		nftKey,
		sdk.NewTransientStoreKey("transient_test"),
	).Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	// gomock initializations
	ctrl := gomock.NewController(suite.T())
	accountKeeper := testutilkeepers.NewMockAccountKeeper(ctrl)
	bankKeeper := testutilkeepers.NewMockBankKeeper(ctrl)

	accountKeeper.EXPECT().GetModuleAddress("nft").Return(suite.addrs[0]).AnyTimes()

	nftKeeper := nftkeeper.NewKeeper(
		nftKey,
		suite.encCfg.Codec,
		accountKeeper,
		bankKeeper,
	)

	obitKeeper := keeper.NewKeeper(suite.encCfg.Codec, key, memKey, nftKeeper)

	suite.obitKeeper = *obitKeeper
	suite.nftKeeper = nftKeeper
	suite.msgServer = keeper.NewMsgServerImpl(*obitKeeper)
	suite.ctx = ctx
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestSaveClass() {
	suite.Assert().False(suite.nftKeeper.HasClass(suite.ctx, class.Id))

	err := suite.obitKeeper.SaveClass(suite.ctx, class)
	suite.Require().NoError(err)

	suite.Assert().True(suite.nftKeeper.HasClass(suite.ctx, class.Id))
}