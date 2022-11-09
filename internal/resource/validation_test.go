package resource

import (
	"testing"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/stretchr/testify/require"
)

func TestValidate_Success(t *testing.T) {
	r := require.New(t)
	r.Nil(Validate(
		cloud.HashicorpCloudLocationLink{
			Location: &cloud.HashicorpCloudLocationLocation{},
			Type:     "resource type",
			ID:       "resource ID",
		},
	))
}

func TestValidate_Error(t *testing.T) {
	tests := map[string]struct {
		input cloud.HashicorpCloudLocationLink
		msg   string
	}{
		"missing resource location fails": {
			input: cloud.HashicorpCloudLocationLink{},
			msg:   "missing resource location",
		},
		"missing resource type fails": {
			input: cloud.HashicorpCloudLocationLink{
				Location: &cloud.HashicorpCloudLocationLocation{},
			},
			msg: "missing resource type",
		},
		"missing resource ID fails": {
			input: cloud.HashicorpCloudLocationLink{
				Location: &cloud.HashicorpCloudLocationLocation{},
				Type:     "resource type",
			},
			msg: "missing resource ID",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)
			r.EqualError(Validate(test.input), test.msg)
		})
	}
}
