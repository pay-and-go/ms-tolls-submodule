FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go

FROM alpine:3 as certs
RUN apk --no-cache add ca-certificates

FROM scratch as app
COPY --from=builder /go/src ./
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./main"]