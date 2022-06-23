FROM golang:1.18-alpine as server-build

RUN mkdir /go/src/goapp

WORKDIR  /go/src/goapp

RUN apk upgrade --update && \
    apk --no-cache add git

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=off

RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/gin-gonic/gin

EXPOSE 18000    


CMD ["realize", "start"]