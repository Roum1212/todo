FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/reminder ./cmd/main.go

FROM alpine

RUN apk update && \
        apk add --no-cache ca-certificates && \
        addgroup -S noroot &&  \
        adduser -S -G noroot noroot && \
        rm -rf /var/cache/apk/*


COPY --from=build /app/bin/reminder /reminder
COPY --from=build /go/bin/goose /usr/local/bin/goose

COPY --from=build /app/migrations /app/migrations

USER noroot

CMD ["/reminder"]
