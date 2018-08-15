FROM golang:1.10.3-alpine3.8

COPY src/api /go/src/api/

WORKDIR /go/src/api/

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

EXPOSE 8080

CMD ["gin", "-i", "run"]
