FROM golang:1.24-bookworm AS base

FROM base AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o hundis

FROM alpine:3.21 AS production

WORKDIR /prod

COPY --from=builder /build/hundis ./

RUN apk add --upgrade --no-cache ca-certificates && update-ca-certificates

EXPOSE 5000

CMD ["/prod/hundis"]