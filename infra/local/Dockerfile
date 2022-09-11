# Example dockerfile for local environment.
FROM golang:alpine

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta
RUN apk update && apk add --no-cache git

WORKDIR /my-ayapingping-go-app

COPY . .

RUN go build -o apiGinGonic src/app/rest/gingonic/gingonic.go

ENTRYPOINT ["/my-ayapingping-go-app/apiGinGonic"]