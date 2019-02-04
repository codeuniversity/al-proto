proto:
	protoc protocol.proto --go_out=plugins=grpc:.

docker-proto:
	docker build -t al-proto-c .
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) al-proto-c
