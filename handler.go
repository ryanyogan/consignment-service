package main

import (
	"context"
	pb "github.com/ryanyogan/consignment-service/proto/consignment"
	vesselProto "github.com/ryanyogan/vessel-service/proto/vessel"
	"log"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - This is the main handler for RPC calls to create
// a new consignment.
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := h.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := h.repository.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments - Runs a collection scan for all consignments
func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := h.repository.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
