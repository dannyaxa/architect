FROM gliderlabs/alpine:3.1

RUN apk-install git go
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:$PATH

RUN go get -u github.com/jteeuwen/go-bindata/...

WORKDIR /go/src/github.com/convox/architect
COPY . /go/src/github.com/convox/architect
RUN go get .
RUN go-bindata -o template.go template/

ENTRYPOINT ["/go/bin/architect"]
