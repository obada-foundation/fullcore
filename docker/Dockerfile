FROM alpine:edge AS builder

RUN apk update && apk upgrade && apk add go

COPY ./src /core

WORKDIR core

RUN go build -o /cored cmd/cored/main.go

FROM alpine:3.14

RUN apk add ca-certificates --no-cache && adduser -D cored

USER cored

WORKDIR /home/cored

COPY --from=builder --chown=cored /cored /home/cored

COPY --chown=cored ./docker/entrypoint.sh /home/cored/

ENTRYPOINT ["/home/cored/entrypoint.sh"]


