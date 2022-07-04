package nodestatus

import (
	"context"
	"runtime"
	"testing"

	requirepkg "github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/hashicorp/hcp-link/pkg/config"
	"github.com/hashicorp/hcp-link/pkg/nodestatus"
)

type stubNodeStatusReporter struct {
	version uint32
	status  proto.Message
}

func (s *stubNodeStatusReporter) GetNodeStatus(_ context.Context) (nodestatus.NodeStatus, error) {
	return nodestatus.NodeStatus{
		StatusVersion: s.version,
		Status:        s.status,
	}, nil
}

func TestNodeStatusService_GetNodeStatus(t *testing.T) {
	require := requirepkg.New(t)

	nodeID := "my-node"
	nodeVersion := "1.2.3"
	statusVersion := uint32(2)
	statusMessage := structpb.NewStringValue("status")

	reporter := &stubNodeStatusReporter{
		version: statusVersion,
		status:  statusMessage,
	}

	service := &Service{
		Config: &config.Config{
			NodeID:             nodeID,
			NodeVersion:        nodeVersion,
			NodeStatusReporter: reporter,
		},
	}

	status, err := service.GetNodeStatus(context.Background(), nil)
	require.NoError(err)

	require.Equal(nodeID, status.NodeId)
	require.Equal(nodeVersion, status.NodeVersion)
	require.Equal(runtime.GOOS, status.NodeOs)
	require.Equal(runtime.GOARCH, status.NodeArchitecture)
	require.NotEmpty(status.Timestamp)
	require.Equal(statusVersion, status.StatusVersion)

	resultStatusMessage := &structpb.Value{}
	err = anypb.UnmarshalTo(status.Status, resultStatusMessage, proto.UnmarshalOptions{})
	require.NoError(err)

	require.Equal(statusMessage.String(), resultStatusMessage.String())
}
