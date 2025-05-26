FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/reminder ./cmd/main.go

FROM alpine

RUN apk update && \
        apk add --no-cache ca-certificates && \
        addgroup -S noroot &&  \
        adduser -S -G noroot noroot && \
        rm -rf /var/cache/apk/*


COPY --from=build /app/bin/reminder /reminder

USER noroot

CMD ["/reminder"]
