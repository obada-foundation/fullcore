FROM golang:1.18-alpine AS builder

RUN apk update && apk add --no-cache curl make git libc-dev bash gcc linux-headers findutils
WORKDIR /go/src/github.com/obada-foundation/fullcore

COPY . .

RUN make vendor && make bin

FROM alpine:3.17

ENV CHAINID "obadatestnet"
ENV DAEMON_HOME "/home/obada/fullcored"
ENV PATH "/home/obada:${PATH}"

RUN apk add ca-certificates --no-cache && adduser -D obada

USER obada

WORKDIR /home/obada

COPY --from=builder --chown=obada /go/src/github.com/obada-foundation/fullcore/build/fullcored /home/obada

COPY --chown=obada ./docker/entrypoint.sh /home/obada/

CMD ["/home/obada/entrypoint.sh"]


