package nodestatus

import (
	"context"
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/hashicorp/hcp-link/gen/proto/go/hashicorp/cloud/hcp_link/node_status/v1"
	"github.com/hashicorp/hcp-link/pkg/config"
)

type Service struct {
	// Config contains all dependencies as well as information about the node
	// Link is running on.
	*config.Config

	pb.UnimplementedNodeStatusServiceServer
}

// GetNodeStatus will be used to regularly fetch the node’s current status.
func (s *Service) GetNodeStatus(ctx context.Context, _ *pb.GetNodeStatusRequest) (*pb.GetNodeStatusResponse, error) {
	// Return an error if s.NodeStatusReporter is not set. This should ideally
	// not happen as the service should only be registered when a node status
	// reporter is available.
	if s.NodeStatusReporter == nil {
		return nil, status.Error(codes.NotFound, "no node status reporter has been registered")
	}

	// Get the node's status
	currentStatus, err := s.NodeStatusReporter.GetNodeStatus(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current node status: %v", err)
	}

	// Marshal the current status into a proto.Any message
	anyStatus, err := anypb.New(currentStatus.Status)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to marshal current status into proto.Any")
	}

	// Collect all information and return them
	return &pb.GetNodeStatusResponse{
		NodeId:           s.NodeID,
		NodeVersion:      s.NodeVersion,
		NodeOs:           runtime.GOOS,
		NodeArchitecture: runtime.GOARCH,
		Timestamp:        timestamppb.Now(),
		StatusVersion:    currentStatus.StatusVersion,
		Status:           anyStatus,
	}, nil
}
