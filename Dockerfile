FROM golang:onbuild
ADD . /go/src/github.com/sibyakin/yarfu
RUN go install github.com/sibyakin/yarfu
RUN mkdir -p ./public
ENTRYPOINT ["/go/bin/yarfu"]
EXPOSE 8080
