FROM golang:1.20

ENV SRC_PATH /go

COPY . $SRC_PATH

WORKDIR $SRC_PATH/src

RUN go mod tidy
