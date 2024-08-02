ARG GO_VERSION=1.22.3

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /rest-ws ./cmd

FROM scratch

COPY .env ./

COPY --from=builder /rest-ws /rest-ws

EXPOSE 5050

ENTRYPOINT ["/rest-ws"]
