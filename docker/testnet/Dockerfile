FROM obada/fullcore:develop

ENV TESTNET_VERSION=0.0.1

COPY --chown=obada:obada docker/testnet/nodes /home/obada/nodes

EXPOSE 26657
EXPOSE 26656
EXPOSE 9090

CMD ["sh", "-c", "./fullcored start --trace --home ./nodes/$TESTNET_VERSION/node0/cored --minimum-gas-prices 1rohi"]
