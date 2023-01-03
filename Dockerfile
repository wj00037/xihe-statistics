FROM golang:latest as BUILDER

# build binary
COPY . /go/src/project/xihe-statistics
RUN cd /go/src/project/xihe-statistics && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

# install timezone file
apk add --no-cache tzdata

COPY  --from=BUILDER /go/src/project/xihe-statistics/xihe-statistics /opt/app

ENTRYPOINT ["/opt/app/xihe-statistics"]