FROM golang:1.18 as builder
WORKDIR /go/src
COPY . /go/src
RUN go build -o webserver webserver-multistage.go
RUN ls
RUN pwd

FROM scratch 
WORKDIR /server
COPY --from=builder /go/src/webserver .
RUN ls
ENTRYPOINT /webserver


