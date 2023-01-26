// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nodestatus

import (
	"context"
	"errors"
	"testing"

	requirepkg "github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/hcp-link/pkg/config"
	"github.com/hashicorp/hcp-link/pkg/nodestatus"
)

type errNodeStatusReporter struct {
	err error
}

func (s *errNodeStatusReporter) GetNodeStatus(_ context.Context) (nodestatus.NodeStatus, error) {
	return nodestatus.NodeStatus{}, s.err
}

func TestCollector_CollectPb_Success(t *testing.T) {
	require := requirepkg.New(t)

	collector := &Collector{
		Config: &config.Config{
			NodeID:      nodeID,
			NodeVersion: nodeVersion,
			NodeStatusReporter: &stubNodeStatusReporter{
				version: statusVersion,
				status:  statusMessage,
			},
		},
	}

	pb, err := collector.CollectPb(bgContext)
	require.Nil(err)
	require.NotNil(pb)
}

func TestCollector_CollectPb_NoReporter(t *testing.T) {
	require := requirepkg.New(t)

	collector := &Collector{
		Config: &config.Config{
			NodeID:      nodeID,
			NodeVersion: nodeVersion,
		},
	}

	_, err := collector.CollectPb(bgContext)
	require.Equal(status.Error(codes.NotFound, "no node status reporter has been registered"), err)
}

func TestCollector_CollectPb_ReporterError(t *testing.T) {
	require := requirepkg.New(t)

	collector := &Collector{
		Config: &config.Config{
			NodeID:      nodeID,
			NodeVersion: nodeVersion,
			NodeStatusReporter: &errNodeStatusReporter{
				err: errors.New("reporter error"),
			},
		},
	}

	_, err := collector.CollectPb(bgContext)
	require.Equal(status.Error(codes.Internal, "failed to get current node status: reporter error"), err)
}

func TestCollector_CollectPb_MarshalFail(t *testing.T) {
	require := requirepkg.New(t)

	collector := &Collector{
		Config: &config.Config{
			NodeID:      nodeID,
			NodeVersion: nodeVersion,
			NodeStatusReporter: &stubNodeStatusReporter{
				version: statusVersion,
				status:  nil,
			},
		},
	}

	_, err := collector.CollectPb(bgContext)
	require.Equal(status.Error(codes.Internal, "failed to marshal current status into proto.Any"), err)
}
