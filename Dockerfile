FROM golang:1.18 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -buildvcs=false -v ./...

FROM golang:1.18
WORKDIR /usr/local/bin
COPY --from=builder /usr/src/app/backend .
CMD ["backend"]
