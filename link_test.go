// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package link

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	scada "github.com/hashicorp/hcp-scada-provider"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	sdk "github.com/hashicorp/hcp-sdk-go/config"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	pb "github.com/hashicorp/hcp-link/gen/proto/go/hashicorp/cloud/hcp_link/node_status/v1"
	nodestatusinternal "github.com/hashicorp/hcp-link/internal/nodestatus"
	"github.com/hashicorp/hcp-link/pkg/config"
	"github.com/hashicorp/hcp-link/pkg/nodestatus"
)

type stubHCPConfig struct {
	sdk.HCPConfig

	apiAddress   string
	apiTLSConfig *tls.Config
}

func (sc stubHCPConfig) APIAddress() string {
	return sc.apiAddress
}

func (sc stubHCPConfig) APITLSConfig() *tls.Config {
	return sc.apiTLSConfig
}

func TestLink(t *testing.T) {
	t.Run("Link library initialization fails if no config is provided", func(t *testing.T) {
		_, err := New(nil)

		r := require.New(t)
		r.EqualError(err, "failed to initialize link library: config must be provided")
	})

	t.Run("Link library initialization succeeds when a valid config is passed", func(t *testing.T) {
		link, err := New(&config.Config{
			NodeID:      "Node ID",
			NodeVersion: "0.0.0",
			Resource: models.HashicorpCloudLocationLink{
				ID:   "ID",
				Type: "Type",
				Location: &models.HashicorpCloudLocationLocation{
					OrganizationID: uuid.New().String(),
					ProjectID:      uuid.New().String(),
				},
			},
			HCPConfig:     stubHCPConfig{},
			SCADAProvider: idleSCADAProvider(),
			Logger:        hclog.Default(),
		})

		r := require.New(t)
		r.NoError(err)
		r.NotNil(link)
	})

	t.Run("Link adds metadata to SCADA provider when started", func(t *testing.T) {
		expectedNodeVersion := "VersionToExpect"
		expectedNodeID := "IDToExpect"
		scadaProvider := idleSCADAProvider()

		givenLink, _ := New(&config.Config{
			NodeID:      expectedNodeID,
			NodeVersion: expectedNodeVersion,
			Resource: models.HashicorpCloudLocationLink{
				ID:   "ID",
				Type: "Type",
				Location: &models.HashicorpCloudLocationLocation{
					OrganizationID: uuid.New().String(),
					ProjectID:      uuid.New().String(),
				},
			},
			HCPConfig:     stubHCPConfig{},
			SCADAProvider: scadaProvider,
			Logger:        hclog.Default(),
		})

		err := givenLink.Start()

		r := require.New(t)
		r.NoError(err)

		r.Equal(expectedNodeID, scadaProvider.GetMeta()["link.node_id"])
		r.Equal(expectedNodeVersion, scadaProvider.GetMeta()["link.node_version"])
	})

	t.Run("Link stops", func(t *testing.T) {
		r := require.New(t)

		scadaProvider := idleSCADAProvider()

		givenLink, _ := New(&config.Config{
			Resource:      models.HashicorpCloudLocationLink{},
			HCPConfig:     stubHCPConfig{},
			SCADAProvider: scadaProvider,
			Logger:        hclog.Default(),
		})

		err := givenLink.Start()
		r.NoError(err)

		err = givenLink.Stop()
		r.NoError(err)
	})
}

func idleSCADAProvider() scada.SCADAProvider {
	scadaProvider, err := scada.New(&scada.Config{
		Service:   "test",
		HCPConfig: struct{ sdk.HCPConfig }{},
		Resource: models.HashicorpCloudLocationLink{
			ID:       "Service-" + uuid.New().String(),
			Type:     "Type",
			Location: &models.HashicorpCloudLocationLocation{},
			UUID:     "",
		},
		Logger:      hclog.Default(),
		TestBackoff: 0,
	})

	if err != nil {
		log.Fatalf("failed to create a new SCADA provider: %v", err)
	}

	return scadaProvider
}

func TestLink_ReportNodeStatus(t *testing.T) {
	r := require.New(t)

	nodeVersion := "1.2.3"
	resourceID := "my-resource"
	resourceType := "hashicorp.example.linked-cluster"
	organizationID := "2f2e3ddf-66d7-4c41-a541-370aa63670f5"
	projectID := "7779414a-785f-4226-94af-030a1e62e154"

	// Prepare a HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		r.Equal(
			"/link/2022-06-04/status/organization/2f2e3ddf-66d7-4c41-a541-370aa63670f5/project/7779414a-785f-4226-94af-030a1e62e154/hashicorp.example.linked-cluster/my-resource/node/my-node",
			req.URL.String(),
		)

		// Parse the body into a proto message
		body, err := io.ReadAll(req.Body)
		r.NoError(err)

		setStatusRequest := &pb.SetNodeStatusRequest{}
		err = proto.Unmarshal(body, setStatusRequest)
		r.NoError(err)

		// Ensure the status message contains all expected information
		r.Equal(nodeVersion, setStatusRequest.NodeStatus.NodeVersion)
		r.Equal(runtime.GOOS, setStatusRequest.NodeStatus.NodeOs)
		r.Equal(runtime.GOARCH, setStatusRequest.NodeStatus.NodeArchitecture)
		r.Equal(uint32(4), setStatusRequest.NodeStatus.StatusVersion)

		// Unmarshal the status and ensure it is as expected
		status := &structpb.Value{}
		err = setStatusRequest.NodeStatus.Status.UnmarshalTo(status)
		r.NoError(err)
		r.Equal("some value", status.GetStringValue())
	}))

	// Prepare a configuration
	c := &config.Config{
		NodeID:      "my-node",
		NodeVersion: "1.2.3",
		Resource: models.HashicorpCloudLocationLink{
			ID:   resourceID,
			Type: resourceType,
			Location: &models.HashicorpCloudLocationLocation{
				OrganizationID: organizationID,
				ProjectID:      projectID,
			},
		},
		HCPConfig: stubHCPConfig{
			apiAddress:   server.Listener.Addr().String(),
			apiTLSConfig: nil,
		},
		// Use a stub node status reporter
		NodeStatusReporter: &nodeStatusReporter{},
	}

	// Create link
	givenLink := link{
		Config:    c,
		apiClient: server.Client(),
		collector: &nodestatusinternal.Collector{
			Config: c,
		},
	}

	// Exercise
	err := givenLink.ReportNodeStatus(context.Background())
	r.NoError(err)
}

type nodeStatusReporter struct{}

func (n *nodeStatusReporter) GetNodeStatus(ctx context.Context) (nodestatus.NodeStatus, error) {
	return nodestatus.NodeStatus{
		StatusVersion: 4,
		Status:        structpb.NewStringValue("some value"),
	}, nil
}
