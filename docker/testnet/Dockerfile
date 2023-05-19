FROM obada/fullcore:develop

ARG TESTNET_VERSION=2

RUN fullcored init testnode --chain-id obada-testnet && wget https://raw.githubusercontent.com/obada-foundation/testnet/main/testnets/testnet-$TESTNET_VERSION/genesis.json -O $HOME/.fullcore/config/genesis.json && fullcored keys --keyring-backend test --keyring-dir ~/.fullcore/keys add testnode && ADDRESS=$(fullcored keys --keyring-backend test --keyring-dir /home/obada/.fullcore/keys show testnode --address) && fullcored add-genesis-account $ADDRESS 1000000000000000000rohi && fullcored gentx testnode 500000000000000000rohi --keyring-backend test --keyring-dir ~/.fullcore/keys --account-number 0 --sequence 0 --chain-id obada-testnet --gas 1000000 --gas-prices 1rohi && fullcored collect-gentxs && fullcored validate-genesis

EXPOSE 26656 26657 26658

CMD ["sh", "-c", "./fullcored start --trace --minimum-gas-prices 1rohi"]
