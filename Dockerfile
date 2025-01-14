FROM golang:latest AS builder

WORKDIR /ssmif_prep

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /ssmif_prep/bin/ssmif_prep ./cmd/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /ssmif_prep/bin/ssmif_prep .

EXPOSE 8080

CMD ["./ssmif_prep"]
