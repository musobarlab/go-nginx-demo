FROM golang:1.10.3-alpine

WORKDIR /go/src/github.com/purwokertodev/go-nginx-demo/service

# Add the source code
ENV SRC_DIR=/go/src/github.com/purwokertodev/go-nginx-demo/service/

ENV BUILD_PACKAGES="git curl"

ADD . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && curl https://glide.sh/get | sh \
  && glide install \
  && apk del $BUILD_PACKAGES \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o my-service .

# Copy environment variable to source dir
COPY .env $SRC_DIR.env

EXPOSE 9000

ENTRYPOINT ["./my-service"]