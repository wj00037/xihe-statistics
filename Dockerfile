FROM golang:latest as BUILDER

# build binary
COPY . /go/src/project/xihe-statistics
RUN cd /go/src/project/xihe-statistics && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest

# install timezone file
RUN apk add --no-cache tzdata

RUN adduser mindspore -u 5000 -D
USER mindspore
WORKDIR /opt/app/

COPY  --from=BUILDER /go/src/project/xihe-statistics/xihe-statistics /opt/app

ENTRYPOINT ["/opt/app/xihe-statistics"]
