package link

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	scada "github.com/hashicorp/hcp-scada-provider"
	sdk "github.com/hashicorp/hcp-sdk-go/config"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"

	"github.com/hashicorp/hcp-link/pkg/config"
)

type stubHCPConfig struct {
	sdk.HCPConfig
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
			ScadaProvider: idleSCADAProvider(),
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

		givenLink, err := New(&config.Config{
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
			ScadaProvider: scadaProvider,
			Logger:        hclog.Default(),
		})

		err = givenLink.Start()

		r := require.New(t)
		r.NoError(err)

		r.Equal(expectedNodeID, scadaProvider.GetMeta()["link.node_id"])
		r.Equal(expectedNodeVersion, scadaProvider.GetMeta()["link.node_version"])
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
