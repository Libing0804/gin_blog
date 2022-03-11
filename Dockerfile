FROM golang:1.15-alpine
WORKDIR  $GOPATH/src/github.com/libing/blog_gin
COPY . $GOPATH/src/github.com/libing/blog_gin
RUN go build .
EXPOSE 8000
ENTRYPOINT ["./blog_gin"]
