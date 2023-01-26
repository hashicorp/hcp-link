// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nodestatus

import (
	"context"

	"google.golang.org/protobuf/types/known/structpb"
)

var (
	nodeID        = "my-node"
	nodeVersion   = "1.2.3"
	statusVersion = uint32(2)
	statusMessage = structpb.NewStringValue("status")
	bgContext     = context.Background()
)
