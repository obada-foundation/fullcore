package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/core/header"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/nft"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"cosmossdk.io/x/nft/module"
	nfttestutil "cosmossdk.io/x/nft/testutil"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/obada-foundation/fullcore/x/obit/keeper"
	obittypes "github.com/obada-foundation/fullcore/x/obit/types"
)

var (
	class = obittypes.Class{
		Id:     "OBT",
		Name:   "test name",
		Symbol: "TST",
		Uri:    "https://obada.io",
	}

	testDID = "did:obada:c7cb2755b93036e2271b487b416a183799733294273ba3c8e2cdab2fc530b005"
)

// KeeperTestSuite is the keeper test suite
type KeeperTestSuite struct {
	suite.Suite

	addrs []sdk.AccAddress
	ctx   sdk.Context

	msgServer   obittypes.MsgServer
	queryServer obittypes.QueryServer

	obitKeeper keeper.Keeper
	nftKeeper  nftkeeper.Keeper

	encCfg moduletestutil.TestEncodingConfig
}

// SetupTest creates a test suite to test the obit module
func (s *KeeperTestSuite) SetupTest() {
	s.encCfg = moduletestutil.MakeTestEncodingConfig(module.AppModuleBasic{})
	nftKey := storetypes.NewKVStoreKey(nft.StoreKey)
	key := storetypes.NewKVStoreKey(obittypes.StoreKey)
	s.addrs = simtestutil.CreateIncrementalAccounts(3)

	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(s.T(), nftKey, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithHeaderInfo(header.Info{Time: time.Now().Round(0).UTC()})

	nftStoreService := runtime.NewKVStoreService(nftKey)

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	accountKeeper := nfttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := nfttestutil.NewMockBankKeeper(ctrl)
	accountKeeper.EXPECT().GetModuleAddress("nft").Return(s.addrs[0]).AnyTimes()
	accountKeeper.EXPECT().AddressCodec().Return(address.NewBech32Codec("obada")).AnyTimes()

	nftKeeper := nftkeeper.NewKeeper(nftStoreService, s.encCfg.Codec, accountKeeper, bankKeeper)
	obitKeeper := keeper.NewKeeper(s.encCfg.Codec, storeService, nftKeeper)

	s.obitKeeper = obitKeeper
	s.nftKeeper = nftKeeper
	s.msgServer = keeper.NewMsgServerImpl(obitKeeper)
	s.queryServer = keeper.NewQueryServerImpl(obitKeeper)
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
