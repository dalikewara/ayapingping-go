FROM golang:alpine

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta
RUN apk update && apk add --no-cache git
RUN apk add --no-cache build-base
RUN apk add --no-cache make

WORKDIR /ayapingping-go

COPY . .

RUN go build -o service main.go

ENTRYPOINT ["/ayapingping-go/service"]