package linkstatus

import (
	"context"
	"testing"

	requirepkg "github.com/stretchr/testify/require"

	"github.com/hashicorp/hcp-link/pkg/config"
	"github.com/hashicorp/hcp-link/pkg/nodestatus"
)

func TestLinkStatusService_GetLinkStatus(t *testing.T) {
	require := requirepkg.New(t)

	nodeID := "my-node"

	service := &Service{
		Config: &config.Config{
			NodeID: nodeID,
			NodeStatusReporter: struct {
				nodestatus.Reporter
			}{},
		},
	}

	status, err := service.GetLinkStatus(context.Background(), nil)
	require.NoError(err)

	require.Equal(nodeID, status.NodeId)
	require.Equal(linkStatusVersion, status.Version)
	require.True(status.Features.NodeStatusReporting.Enabled)
}

func TestLinkStatusService_GetLinkStatus_NoNodeStatusReporter(t *testing.T) {
	require := requirepkg.New(t)

	nodeID := "my-node"

	service := &Service{
		Config: &config.Config{
			NodeID:             nodeID,
			NodeStatusReporter: nil,
		},
	}

	status, err := service.GetLinkStatus(context.Background(), nil)
	require.NoError(err)

	require.Equal(nodeID, status.NodeId)
	require.Equal(linkStatusVersion, status.Version)
	require.False(status.Features.NodeStatusReporting.Enabled)
}
