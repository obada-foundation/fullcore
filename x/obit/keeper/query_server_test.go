package keeper_test

import "github.com/obada-foundation/fullcore/x/obit/types"

// TestMintNFT tests that a user can mint an NFT
func (s *KeeperTestSuite) TestGetNFT() {
	s.TestMintNFT()

	resp, err := s.queryServer.GetNFT(s.ctx, &types.QueryGetNFTRequest{
		Id: testDID,
	})
	s.Require().NoError(err)
	s.Equal(testDID, resp.GetNft().Id)
}

// TestGetNFTByAddress tests that a user can get NFTs by address
func (s *KeeperTestSuite) TestGetNFTByAddress() {
	s.TestMintNFT()

	resp, err := s.queryServer.GetNFTByAddress(s.ctx, &types.QueryGetNFTByAddressRequest{
		Address: s.addrs[0].String(),
	})
	s.Require().NoError(err)

	NFTs := resp.GetNft()

	s.Equal(1, len(NFTs))
	s.Equal(testDID, NFTs[0].GetId())
}
