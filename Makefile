proto:
	protoc protocol.proto --go_out=plugins=grpc:.

docker-proto:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) al-proto-c
