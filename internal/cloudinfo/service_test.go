// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudinfo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/banzaicloud/cloudinfo/internal/cloudinfo/types"
)

func TestServiceService_ListServices(t *testing.T) {
	store := NewInMemoryServiceStore()
	store.services = map[string][]types.Service{
		"amazon": {
			{
				Service: "compute",
			},
			{
				Service: "eks",
			},
		},
		"google": {
			{
				Service: "compute",
			},
			{
				Service: "gke",
			},
		},
	}

	serviceService := NewServiceService(store)

	services, err := serviceService.ListServices(context.Background(), "amazon")
	require.NoError(t, err)

	assert.Equal(
		t,
		[]Service{
			{Code: "compute", providerName: "amazon"}, {Code: "eks", providerName: "amazon"},
		},
		services,
	)
}
