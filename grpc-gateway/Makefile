repo=test
tag=dev

all: build push
build:
	docker build --no-cache -t ${repo}/grpc-gateway:${tag} .
push:
	docker push ${repo}/grpc-gateway:${tag}
