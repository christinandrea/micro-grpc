package impl

import (
	"internal/protoGen"

	"context"
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
	return &protoGen.AddedRepoResponse{
		AddedRepoResponse: in,
		Error:             nil,
	}, nil
}
