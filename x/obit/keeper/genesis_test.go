package keeper_test

import (
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// TestInitGenesis test init genesis
func (s *KeeperTestSuite) TestInitGenesis() {
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

	s.obitKeeper.InitGenesis(s.ctx, genesisState)
	got := s.obitKeeper.ExportGenesis(s.ctx)

	s.NotNil(got)

	s.Equal(genesisState, *got)
}
