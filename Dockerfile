# Build the Akasha gateway in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go/src/github.com/karalabe/akasha-gateway
RUN go install -v github.com/karalabe/akasha-gateway

# Pull the Akasha gateway into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/akasha-gateway /usr/local/bin/

EXPOSE 8080 30303 30303/udp
ENTRYPOINT ["akasha-gateway"]
