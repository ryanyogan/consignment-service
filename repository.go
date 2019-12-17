package main

import (
	"context"
	pb "github.com/ryanyogan/consignment-service/proto/consignment"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository - defines the repo for MongoDB
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cursor, err := repository.collection.Find(context.Background(), nil, nil)
	var consignments []*pb.Consignment
	for cursor.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cursor.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
