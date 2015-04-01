FROM gliderlabs/alpine:3.1

RUN apk-install go
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:$PATH

WORKDIR /go/src/github.com/convox/architect
COPY . /go/src/github.com/convox/architect
RUN go get .

ENTRYPOINT ["/go/bin/architect"]
