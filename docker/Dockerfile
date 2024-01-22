FROM golang:1.21-alpine AS builder

RUN <<EOF
apk update
apk add --no-cache curl make git libc-dev bash gcc linux-headers findutils
EOF

WORKDIR /go/src/github.com/obada-foundation/fullcore

COPY . .

RUN make src/vendor && make bin

FROM alpine:3.17 as runner

ENV CHAINID "obadatestnet"
ENV DAEMON_HOME "/home/obada/fullcored"
ENV PATH "/home/obada:${PATH}"

RUN apk add ca-certificates bash --no-cache && adduser -D obada

USER obada

WORKDIR /home/obada

COPY --from=builder --chown=obada /go/src/github.com/obada-foundation/fullcore/build/fullcored /home/obada

COPY --chown=obada ./docker/entrypoint.sh /home/obada/
COPY --chown=obada docker/config/app.toml /home/obada/.fullcored/config/app.toml

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

EXPOSE 26656 26657 26660 1317 9090 9091

CMD ["/home/obada/entrypoint.sh"]


