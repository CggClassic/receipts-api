FROM golang:1.21.5 AS builder

WORKDIR /go/src/app/
COPY . .
RUN CGO_ENABLED=0 go build

FROM alpine:3.19 AS runner

COPY --from=builder /go/src/app/receipts-api /app/receipts-api
COPY --from=builder --chmod=755 /go/src/app/index.html /app/index.html
RUN chmod +x /app/receipts-api
ENTRYPOINT ["/app/receipts-api"]