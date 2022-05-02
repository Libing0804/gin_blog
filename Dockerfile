FROM golang:alpine
WORKDIR  $GOPATH/src/github.com/libing/blog_gin
COPY . $GOPATH/src/github.com/libing/blog_gin
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
EXPOSE 8000
ENTRYPOINT ["go","mod","tidy"]
ENTRYPOINT ["go","run","main.go"]
