# Build janeta in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make git

ADD . /go/src/github.com/janotchain/janeta
RUN cd /go/src/github.com/janotchain/janeta && make janetad && make janetacli

# Pull janeta into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/janotchain/janeta/cmd/janetad/janetad /usr/local/bin/
COPY --from=builder /go/src/github.com/janotchain/janeta/cmd/janetacli/janetacli /usr/local/bin/

EXPOSE 1999 46656 46657 9888
