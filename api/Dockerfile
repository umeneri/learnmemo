FROM golang:1.12.0 as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN groupadd -g 10001 myapp \
    && useradd -u 10001 -g myapp myapp

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

# FROM scratch
FROM alpine:latest

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /go/src/app/public /go/src/app/public
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /go/src/app

EXPOSE 8080

USER myapp

ENTRYPOINT ["/go/bin/app"]