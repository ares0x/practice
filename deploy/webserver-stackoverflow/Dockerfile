FROM golang:1.20-alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server

FROM scratch
COPY --from=builder /app/server /opt/app/ 

EXPOSE 8080 

CMD ["/opt/app/server"]


