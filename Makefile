build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consignment-service .
	docker build -t transport-service-consignment .

run:
	docker run -p  50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		transport-service-consignment