FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o fee_tracker cmd/fee_tracker/main.go

FROM golang:1.23.1

WORKDIR /app

COPY --from=builder /app/fee_tracker .

CMD ["./fee_tracker"]