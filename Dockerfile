ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN go build -o /paloma cmd/app/main.go


FROM alpine:latest
COPY --from=builder /paloma /usr/local/bin/

# goose binary
COPY --from=builder /go/bin/goose /usr/local/bin/

# migrations files
COPY --from=builder /usr/src/app/db/migrations /migrations

CMD ["paloma"]
