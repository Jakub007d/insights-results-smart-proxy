// Copyright 2021 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"fmt"

	"github.com/RedHatInsights/insights-results-smart-proxy/amsclient"
	"github.com/RedHatInsights/insights-results-smart-proxy/types"
)

type mockAMSClient struct {
	clustersPerOrg map[types.OrgID][]types.ClusterInfo
}

func (m *mockAMSClient) GetClustersForOrganization(
	orgID types.OrgID,
	unused1, unused2 []string,
) (
	clusterInfoList []types.ClusterInfo,
	clusterNamesMap map[types.ClusterName]string,
	err error,
) {

	clusterInfoList, ok := m.clustersPerOrg[orgID]
	if !ok {
		return nil, nil, fmt.Errorf("No clusters")
	}

	clusterNamesMap = make(map[types.ClusterName]string)
	for i := range clusterInfoList {
		clusterNamesMap[clusterInfoList[i].ID] = clusterInfoList[i].DisplayName
	}

	return
}

// AMSClientWithOrgResults creates a mock of AMSClient interface that returns the results
// defined by orgID and clusters parameters
func AMSClientWithOrgResults(orgID types.OrgID, clusters []types.ClusterInfo) amsclient.AMSClient {
	return &mockAMSClient{
		clustersPerOrg: map[types.OrgID][]types.ClusterInfo{
			orgID: clusters,
		},
	}
}
