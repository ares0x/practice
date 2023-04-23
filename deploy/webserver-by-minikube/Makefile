export tag=v0.4
author = ares0x
servername = webserver

KUBECTL := kubectl

root:
	export root=github.com/ares0x/practice

build:
	echo "building webserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64

release:
	echo "building webserver container"
	docker build -t $(author)/$(servername):${tag} .

unuse:
	eval $(minikube docker-env -u)

use:
	eval $(minikube docker-env)
