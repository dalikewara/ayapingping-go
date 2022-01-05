# Your dev dockerfile here...
# Example:
FROM golang:alpine

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta
RUN apk update && apk add --no-cache git

WORKDIR /my-ayapingping-go-app

COPY . .

RUN go build -o api src/apps/api/main.go

ENTRYPOINT ["/my-ayapingping-go-app/api"]