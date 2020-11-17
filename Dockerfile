FROM golang:alpine AS builder

ENV GO111MODULE="on" \
    CGO_ENABLED=1 \
    GOOS=linux

RUN apk --no-cache add build-base
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go generate ./ent --wipe \
 && go build -ldflags="-s -w" -o /var/app/main cmd/app/main.go

FROM alpine
RUN apk --no-cache add ca-certificates

COPY --from=builder /var/app/main /var/app/

EXPOSE 8080
ENTRYPOINT [ "/var/app/main" ]
