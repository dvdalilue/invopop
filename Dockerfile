FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go build -o server cmd/server/main.go
RUN go build -o client cmd/client/main.go

FROM alpine

RUN adduser -S -D -H -h /home/app appuser

USER appuser

WORKDIR /home/app

COPY --from=builder /build/server .
COPY --from=builder /build/client .
CMD ["./server"]