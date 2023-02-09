package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	module "github.com/obada-foundation/fullcore/x/obit"
	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	testutilctx "github.com/obada-foundation/fullcore/testutil/context"
	testutilcodec "github.com/obada-foundation/fullcore/testutil/module/codec"
	testutil "github.com/obada-foundation/fullcore/x/obit/testutil"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx sdk.Context

	msgServer  types.MsgServer
	obitKeeper keeper.Keeper

	encCfg testutilcodec.TestEncodingConfig
}

func (suite *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	memKey := sdk.NewKVStoreKey(types.MemStoreKey)

	suite.encCfg = testutilcodec.MakeTestEncodingConfig(module.AppModuleBasic{})

	testCtx := testutilctx.DefaultContextWithDB(
		suite.T(),
		key,
		sdk.NewTransientStoreKey("transient_test"),
	)
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	// gomock initializations
	ctrl := gomock.NewController(suite.T())
	nftKeeper := testutil.NewMockNftKeeper(ctrl)
	obitKeeper := keeper.NewKeeper(suite.encCfg.Codec, key, memKey, nftKeeper)

	suite.obitKeeper = *obitKeeper
	suite.ctx = ctx

}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestSaveClass() {
	suite.obitKeeper.SaveClass(suite.ctx, types.Class{
		Id:     "test",
		Name:   "test name",
		Symbol: "TST",
		Uri:    "https://obada.io",
	})
}
