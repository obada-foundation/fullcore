package keeper_test

import (
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (suite *KeeperTestSuite) TestInitGenesis() {
	genesisState := types.GenesisState{
		Classes: []*types.Class{
			{
				Id:     types.OBTClass,
				Name:   "Obada network NFT Token",
				Symbol: types.OBTClass,
				Uri:    "https://obada.io",
			},
		},
	}

	suite.obitKeeper.InitGenesis(suite.ctx, genesisState)
	got := suite.obitKeeper.ExportGenesis(suite.ctx)

	suite.NotNil(got)

	suite.Equal(genesisState, *got)
}
