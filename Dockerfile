FROM golang:1.15.14-alpine3.14 as build
ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED 0
WORKDIR $GOPATH/src/jellydelete
COPY . $GOPATH/src/jellydelete
RUN go build .


FROM alpine:3.14
WORKDIR /usr/src/app
COPY --from=build /go/src/jellydelete/jellydelete  .
EXPOSE 8888
ENTRYPOINT ["./jellydelete"]

# 未将configs写入镜像，单独指定、或者configmap;