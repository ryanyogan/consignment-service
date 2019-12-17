package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/ryanyogan/consignment-service/proto/consignment"
	vesselProto "github.com/ryanyogan/vessel-service/proto/vessel"
)

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("transport.service.consignment"),
	)

	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("transport").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("transport.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
