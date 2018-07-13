FROM golang:1.9-alpine
RUN mkdir -p $GOPATH/src/github.com/chrislusf/seaweedfs
COPY . $GOPATH/src/github.com/chrislusf/seaweedfs
WORKDIR $GOPATH/src/github.com/chrislusf/seaweedfs
RUN go build -o weed/weed ./weed/
EXPOSE 8080
EXPOSE 9333
VOLUME /data
ENTRYPOINT ["weed/weed"]