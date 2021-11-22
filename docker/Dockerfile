FROM alpine:edge AS builder

RUN apk update && apk upgrade && apk add go

COPY ./src /fullcore

WORKDIR fullcore

RUN go build -o /fullcored cmd/fullcored/main.go

FROM alpine:3.14

ENV CHAINID "obadatestnet"
ENV DAEMON_HOME "/home/obada/fullcored"
ENV PATH "/home/obada:${PATH}"

RUN apk add ca-certificates --no-cache && adduser -D obada

USER obada

WORKDIR /home/obada

COPY --from=builder --chown=obada /fullcored /home/obada

COPY --chown=obada ./docker/entrypoint.sh /home/obada/

CMD ["/home/cored/entrypoint.sh"]


