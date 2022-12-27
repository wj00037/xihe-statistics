FROM golang:latest as BUILDER

# build binary
COPY . /go/src/project/xihe-statistics
RUN cd /go/src/project/xihe-statistics && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

COPY  --from=BUILDER /go/src/project/xihe-statistics/xihe-statistics /opt/app

ENTRYPOINT ["/opt/app/xihe-statistics"]