// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nodestatus

import (
	"context"
	"runtime"
	"testing"

	requirepkg "github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
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

func TestNodeStatusService_GetNodeStatus_Success(t *testing.T) {
	require := requirepkg.New(t)

	reporter := &stubNodeStatusReporter{
		version: statusVersion,
		status:  statusMessage,
	}

	service := &Service{
		Collector: &Collector{
			Config: &config.Config{
				NodeID:             nodeID,
				NodeVersion:        nodeVersion,
				NodeStatusReporter: reporter,
			},
		},
	}

	status, err := service.GetNodeStatus(bgContext, nil)
	require.NoError(err)

	nodeStatus := status.NodeStatus

	require.Equal(nodeID, nodeStatus.NodeId)
	require.Equal(nodeVersion, nodeStatus.NodeVersion)
	require.Equal(runtime.GOOS, nodeStatus.NodeOs)
	require.Equal(runtime.GOARCH, nodeStatus.NodeArchitecture)
	require.NotEmpty(nodeStatus.Timestamp)
	require.Equal(statusVersion, nodeStatus.StatusVersion)

	resultStatusMessage := &structpb.Value{}
	err = anypb.UnmarshalTo(nodeStatus.Status, resultStatusMessage, proto.UnmarshalOptions{})
	require.NoError(err)

	require.Equal(statusMessage.String(), resultStatusMessage.String())

}

func TestNodeStatusService_GetNodeStatus_Error(t *testing.T) {
	require := requirepkg.New(t)

	service := &Service{
		Collector: &Collector{
			Config: &config.Config{
				NodeID:             nodeID,
				NodeVersion:        nodeVersion,
				NodeStatusReporter: nil,
			},
		},
	}

	_, err := service.GetNodeStatus(bgContext, nil)

	require.Equal(grpcstatus.Error(codes.NotFound, "no node status reporter has been registered"), err)
}
