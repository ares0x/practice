FROM golang:1.18

WORKDIR /go/src

COPY . /go/src

RUN go build -o /webserver webserver.go
# 设置对外暴露的端口
EXPOSE 8080

ENTRYPOINT ["/webserver"]