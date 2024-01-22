FROM obada/fullcore:develop

ARG TESTNET_VERSION=2

RUN <<EOF
fullcored init testnode --chain-id obada-testnet --default-denom rohi
wget https://raw.githubusercontent.com/obada-foundation/testnet/main/testnets/testnet-$TESTNET_VERSION/genesis.json -O $HOME/.fullcored/config/genesis.json 
fullcored keys --keyring-backend test --keyring-dir ~/.fullcored/keys add testnode 
ADDRESS=$(fullcored keys --keyring-backend test --keyring-dir /home/obada/.fullcored/keys show testnode --address) 
fullcored genesis add-genesis-account $ADDRESS 1000000000000000000rohi

fullcored genesis gentx testnode 500000000000000000rohi \
    --keyring-backend test \
    --keyring-dir ~/.fullcored/keys \
    --account-number 0 \
    --sequence 0 \
    --chain-id obada-testnet \
    --gas 1000000 \
    --gas-prices 1rohi

fullcored genesis collect-gentxs
fullcored genesis validate-genesis
EOF

EXPOSE 26656 26657 26658

CMD ["sh", "-c", "./fullcored start --trace --minimum-gas-prices 1rohi"]
