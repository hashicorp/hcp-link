package linkstatus

import (
	"context"

	pb "github.com/hashicorp/hcp-link/gen/proto/go/hashicorp/cloud/hcp_link/link_status/v1"
	"github.com/hashicorp/hcp-link/pkg/config"
)

const (
	linkStatusVersion = "0.0.1"
)

type Service struct {
	// Config contains all dependencies as well as information about the node
	// Link is running on.
	*config.Config

	pb.UnimplementedLinkStatusServiceServer
}

// GetLinkStatus will be used to fetch the node’s link specific status.
func (s *Service) GetLinkStatus(ctx context.Context, _ *pb.GetLinkStatusRequest) (*pb.GetLinkStatusResponse, error) {
	return &pb.GetLinkStatusResponse{
		NodeId:  s.NodeID,
		Version: linkStatusVersion,
	}, nil
}
