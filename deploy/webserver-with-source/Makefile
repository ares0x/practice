export tag=v0.2

root:
	export root=github.com/ares0x/practice

build:
	echo "building webserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64

release:
	echo "building webserver container"
	docker build -t ares0x/webserver:${tag} .