Dockerfile.lib.matador.ais.co.thARG GO_VERSION=1.22-alpine3.20
FROM golang:${GO_VERSION} AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
ENV TZ=Asia/Bangkok

EXPOSE 3000
CMD ["/app/main"]