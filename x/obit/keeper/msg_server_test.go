package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (suite *KeeperTestSuite) TestMintNFT() {
	creator := suite.addrs[0]

	err := suite.obitKeeper.SaveClass(suite.ctx, class)
	suite.Require().NoError(err)

	msg := types.MsgMintNFT{
		Creator: creator.String(),
		Id:      "did:obada:c7cb2755b93036e2271b487b416a183799733294273ba3c8e2cdab2fc530b005",
		Usn:     "2z41UH7n",
	}

	resp, err := suite.msgServer.MintNFT(suite.ctx, &msg)
	suite.Require().NoError(err)

	suite.Assert().Equal(msg.Id, resp.Id)

	// Test that we cannot mint it a second time
	resp, err = suite.msgServer.MintNFT(suite.ctx, &msg)
	suite.Require().ErrorIs(err, nft.ErrNFTExists)

	state := suite.nftKeeper.ExportGenesis(suite.ctx)
	suite.Require().Len(state.Entries, 1)

	nft, err := suite.obitKeeper.GetNFT(suite.ctx, &types.QueryGetNFTRequest{
		Id: msg.Id,
	})
	suite.Require().NoError(err)

	suite.Equal(msg.Id, nft.Id)
	suite.Equal(types.OBTClass, nft.ClassId)

	nfts, err := suite.obitKeeper.GetNFTByAddress(suite.ctx, &types.QueryGetNFTByAddressRequest{
		Address: msg.Creator,
	})
	suite.Require().NoError(err)
	suite.Require().Len(nfts.NFT, 1)
}

func (suite *KeeperTestSuite) TestTransferNFT() {
	creator := suite.addrs[0]

	err := suite.obitKeeper.SaveClass(suite.ctx, class)
	suite.Require().NoError(err)

	msg := types.MsgMintNFT{
		Creator: creator.String(),
		Id:      "did:obada:c7cb2755b93036e2271b487b416a183799733294273ba3c8e2cdab2fc530b005",
		Usn:     "2z41UH7n",
	}

	_, err = suite.msgServer.MintNFT(suite.ctx, &msg)
	suite.Require().NoError(err)

	_, err = suite.msgServer.TransferNFT(suite.ctx, &types.MsgTransferNFT{
		Id:       msg.Id,
		Sender:   creator.String(),
		Receiver: creator.String(),
	})
	suite.Require().ErrorIs(sdkerrors.ErrInvalidAddress, err)

	newOwner := suite.addrs[1].String()

	_, err = suite.msgServer.TransferNFT(suite.ctx, &types.MsgTransferNFT{
		Id:       msg.Id,
		Sender:   creator.String(),
		Receiver: newOwner,
	})
	suite.Require().NoError(err)

	nfts, err := suite.obitKeeper.GetNFTByAddress(suite.ctx, &types.QueryGetNFTByAddressRequest{
		Address: msg.Creator,
	})
	suite.Require().NoError(err)
	suite.Require().Len(nfts.NFT, 0)

	nfts, err = suite.obitKeeper.GetNFTByAddress(suite.ctx, &types.QueryGetNFTByAddressRequest{
		Address: newOwner,
	})
	suite.Require().NoError(err)
	suite.Require().Len(nfts.NFT, 1)
}
