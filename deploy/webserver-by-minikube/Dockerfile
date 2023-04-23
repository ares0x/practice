FROM golang:1.18-alpine as builder
WORKDIR /usr/src/app
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build  -ldflags "-s -w" -o webserver

FROM scratch
COPY --from=builder /usr/src/app/webserver /opt/app/

EXPOSE 8888

CMD ["/opt/app/webserver"]

