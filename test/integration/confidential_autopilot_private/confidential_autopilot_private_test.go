// Copyright 2025 Google LLC
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

package confidential_autopilot_private

import (
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-google-modules/terraform-google-kubernetes-engine/test/integration/testutils"
)

func TestConfidentialAutopilotPrivate(t *testing.T) {
	projectID := testutils.GetTestProjectFromSetup(t, 1)
	bpt := tft.NewTFBlueprintTest(t,
		tft.WithVars(map[string]interface{}{"project_id": projectID}),
		tft.WithRetryableTerraformErrors(testutils.RetryableTransientErrors, 3, 2*time.Minute),
	)

	bpt.DefineVerify(func(assert *assert.Assertions) {
		testutils.TGKEVerify(t, bpt, assert)

		location := bpt.GetStringOutput("location")
		clusterName := bpt.GetStringOutput("cluster_name")
		key := bpt.GetStringOutput("kms_key")

		op := gcloud.Runf(t, "container clusters describe %s --zone %s --project %s", clusterName, location, projectID)
		assert.True(op.Get("autopilot.enabled").Bool(), "should be autopilot")
		assert.Equal(op.Get("autoscaling.autoprovisioningNodePoolDefaults.bootDiskKmsKey").String(), key, "should have CMEK configured in boot disk")
		assert.True(op.Get("confidentialNodes.enabled").Bool(), "should have confidential nodes enabled")
		assert.Equal(op.Get("databaseEncryption.state").String(), "ENCRYPTED", "should have database encryption")
		assert.Equal(op.Get("databaseEncryption.keyName").String(), key, "should have CMEK configured in database")
	})
	bpt.Test()
}
