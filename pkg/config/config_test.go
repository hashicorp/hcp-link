package config

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	scada "github.com/hashicorp/hcp-scada-provider"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	sdk "github.com/hashicorp/hcp-sdk-go/config"
	requirepkg "github.com/stretchr/testify/require"
)

type stubHCPConfig struct {
	sdk.HCPConfig
}

type stubSCADAProvider struct {
	scada.SCADAProvider
}

func validConfig() *Config {
	return &Config{
		NodeID:      "node-id",
		NodeVersion: "0.0.1",
		Resource: cloud.HashicorpCloudLocationLink{
			ID:   "resource-id",
			Type: "hashicorp.test.resource",
			Location: &cloud.HashicorpCloudLocationLocation{
				OrganizationID: "575f732d-b7c5-4df2-85df-8594a114c8e1",
				ProjectID:      "c0a38947-7c11-4038-9cf6-2657a2c67cac",
			},
		},
		HCPConfig:     stubHCPConfig{},
		ScadaProvider: &stubSCADAProvider{},
		Logger:        hclog.NewNullLogger(),
	}
}

func TestConfig_Valid(t *testing.T) {
	require := requirepkg.New(t)
	require.NoError(validConfig().Validate())
}

func TestConfig_Invalid(t *testing.T) {
	testCases := []struct {
		name          string
		mutate        func(*Config)
		expectedError string
	}{
		{
			name: "missing node ID",
			mutate: func(config *Config) {
				config.NodeID = ""
			},
			expectedError: "node ID must be provided",
		},
		{
			name: "missing node version",
			mutate: func(config *Config) {
				config.NodeVersion = ""
			},
			expectedError: "node version must be provided",
		},
		{
			name: "invalid resource",
			mutate: func(config *Config) {
				config.Resource.ID = ""
			},
			expectedError: "resource link is invalid: missing resource ID",
		},
		{
			name: "missing HCP Config",
			mutate: func(config *Config) {
				config.HCPConfig = nil
			},
			expectedError: "HCP config must be provided",
		},
		{
			name: "missing SCADA Provider",
			mutate: func(config *Config) {
				config.ScadaProvider = nil
			},
			expectedError: "SCADA provider must be provided",
		},
		{
			name: "missing Logger",
			mutate: func(config *Config) {
				config.Logger = nil
			},
			expectedError: "logger must be provided",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require := requirepkg.New(t)

			config := validConfig()
			testCase.mutate(config)

			require.EqualError(config.Validate(), testCase.expectedError)
		})
	}
}
