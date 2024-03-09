FROM golang:1.19-alpine

# Install system dependecies
RUN apk update
RUN apk add --no-cache git
RUN apk add --no-cache tzdata
RUN apk add --no-cache build-base
RUN apk add --no-cache make
ENV TZ Asia/Jakarta

# Setup the app
WORKDIR /ayapingping-go
COPY . .
RUN chmod +x infra/chmod.sh
RUN ./infra/chmod.sh

# Build the app
RUN ./infra/docker-build.sh

# Run apps
ENTRYPOINT ["/ayapingping-go/infra/docker-entrypoint.sh"]