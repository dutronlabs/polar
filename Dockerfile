FROM golang:1.5-alpine
MAINTAINER chris dutra
RUN apk update; apk add git
RUN go get -u github.com/coreos/clair/contrib/analyze-local-images
RUN go get -u github.com/dutronlabs/polar
EXPOSE 9001
CMD ["polar"]