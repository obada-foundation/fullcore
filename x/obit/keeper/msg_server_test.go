package keeper_test

import (
	"cosmossdk.io/x/nft"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// TestMintNFT tests that a user can mint an NFT
func (s *KeeperTestSuite) TestMintNFT() {
	creator := s.addrs[0]

	err := s.obitKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	msg := types.MsgMintNFT{
		Creator: creator.String(),
		Id:      testDID,
		Usn:     "2z41UH7n",
	}

	resp, err := s.msgServer.MintNFT(s.ctx, &msg)
	s.Require().NoError(err)

	s.Assert().Equal(msg.Id, resp.Id)

	// Test that we cannot mint it a second time
	_, err = s.msgServer.MintNFT(s.ctx, &msg)
	s.Require().ErrorIs(err, nft.ErrNFTExists)

	state := s.nftKeeper.ExportGenesis(s.ctx)
	s.Require().Len(state.Entries, 1)

	nftResp, err := s.queryServer.GetNFT(s.ctx, &types.QueryGetNFTRequest{
		Id: msg.Id,
	})
	s.Require().NoError(err)

	NFT := nftResp.GetNft()

	s.Equal(msg.Id, NFT.Id)
	s.Equal(types.OBTClass, NFT.ClassId)

	nfts, err := s.queryServer.GetNFTByAddress(s.ctx, &types.QueryGetNFTByAddressRequest{
		Address: msg.Creator,
	})
	s.Require().NoError(err)
	s.Require().Len(nfts.GetNft(), 1)
}

// TestTransferNFT tests that a user can transfer an NFT
func (s *KeeperTestSuite) TestTransferNFT() {
	creator := s.addrs[0]

	err := s.obitKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	msg := types.MsgMintNFT{
		Creator: creator.String(),
		Id:      "did:obada:c7cb2755b93036e2271b487b416a183799733294273ba3c8e2cdab2fc530b005",
		Usn:     "2z41UH7n",
	}

	_, err = s.msgServer.MintNFT(s.ctx, &msg)
	s.Require().NoError(err)

	_, err = s.msgServer.TransferNFT(s.ctx, &types.MsgTransferNFT{
		Id:       msg.Id,
		Sender:   creator.String(),
		Receiver: creator.String(),
	})
	s.Require().ErrorIs(sdkerrors.ErrInvalidAddress, err)

	newOwner := s.addrs[1].String()

	_, err = s.msgServer.TransferNFT(s.ctx, &types.MsgTransferNFT{
		Id:       msg.Id,
		Sender:   creator.String(),
		Receiver: newOwner,
	})
	s.Require().NoError(err)

	nfts, err := s.queryServer.GetNFTByAddress(s.ctx, &types.QueryGetNFTByAddressRequest{
		Address: msg.Creator,
	})
	s.Require().NoError(err)
	s.Require().Len(nfts.GetNft(), 0)

	nfts, err = s.queryServer.GetNFTByAddress(s.ctx, &types.QueryGetNFTByAddressRequest{
		Address: newOwner,
	})
	s.Require().NoError(err)
	s.Require().Len(nfts.GetNft(), 1)
}
