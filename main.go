package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	pb "github.com/ryanyogan/consignment-service/proto/consignment"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repo, this simulates the use of a datastore
// of some kind. We'll replace with real later
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satifisy the service
// we defined in our protobuf definition.  You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments - will return all consignments from the active repository
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("transport.service.consignment"),
	)

	srv.Init()

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
