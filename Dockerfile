FROM golang:1.18.5 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -buildvcs=false -o /usr/local/bin -v ./...

FROM golang:1.18.5
WORKDIR /usr/local/bin
COPY --from=builder /usr/local/bin/backend .
CMD ["backend", "migrate"]
