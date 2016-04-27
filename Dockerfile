FROM golang:1.5-alpine
MAINTAINER chris dutra
RUN apk update; apk add git
RUN go get -u github.com/coreos/clair/contrib/analyze-local-images; go get github.com/gin-gonic/gin; go get gopkg.in/redis.v3
COPY . $GOPATH/src/github.com/dutronlabs/polar
RUN go install github.com/dutronlabs/polar
RUN ls $GOPATH/bin
EXPOSE 9001
CMD ["polar"]