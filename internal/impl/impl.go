package impl

import (
	protoGen "github.com/christinandrea/micro-grpc/internal/protoGen"

	"context"
	"log"
)

// RepositoryImpl struct to implement repository
type RepositoryImpl struct {
}

//NewRepoImpl returns pointer for implementation
func NewRepoImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}

//Add function to implement gRPC service
func (repoImpl *RepositoryImpl) Add(ctx context.Context, in *protoGen.Repository) (*protoGen.AddedRepoResponse, error) {

	log.Println("Repository persisted to the storage")

	return &protoGen.AddedRepoResponse{
		AddedRepo: in,
		Error:     nil,
	}, nil
}
