ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o ./chucklerama ./cmd/main.go
CMD ["ls", "-l"]

FROM debian:bookworm

COPY --from=builder /usr/src/app/chucklerama /usr/local/bin/
EXPOSE 8080
CMD ["chucklerama"]
