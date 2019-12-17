build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consignment-service .
	docker build -t transport-service-consignment .

run:
	docker run -d -p  50051:50051 transport-service-consignment