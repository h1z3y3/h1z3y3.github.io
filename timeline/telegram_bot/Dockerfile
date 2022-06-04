FROM r.so.qihoo.net/library/golang:1.17.8-alpine AS builder

COPY . /go/src/app
WORKDIR /go/src/app

RUN go build -v -o app . && chmod a+x app

FROM r.so.qihoo.net/library/alpine:3.11

COPY --from=builder /go/src/app/app .

ENTRYPOINT ["./app"]
