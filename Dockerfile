FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go build -o server cmd/server/main.go

FROM alpine

RUN adduser -S -D -H -h /home/app appuser

USER appuser

WORKDIR /home/app

COPY --from=builder /build/server .
CMD ["./server"]