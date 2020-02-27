# MAINTAINER Jonathan Mainguy <jon@soh.re>
FROM golang:1.13.7
WORKDIR /go/src/app
ENV GO111MODULE=on
ADD . .
RUN go build
CMD ["/go/src/app/run.sh"]
