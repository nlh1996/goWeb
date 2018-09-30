FROM golang
MAINTAINER nlh1996 "nlh.lxz@foxmail.com"

ADD . /go/src/
WORKDIR /go/

RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2 
RUN go install /go/src/demo

ENTRYPOINT /go/bin/demo

EXPOSE 8000

