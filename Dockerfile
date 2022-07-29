FROM golang:1.18.4 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -buildvcs=false -o /usr/local/bin -v ./...

FROM golang:1.18.4
WORKDIR /usr/local/bin
COPY --from=builder /usr/local/bin/backend .
CMD ["backend"]
