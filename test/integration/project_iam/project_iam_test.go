/**
 * Copyright 2026 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package project_iam

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

func TestProjectIam(t *testing.T) {
	test := tft.NewTFBlueprintTest(t)

	test.DefineVerify(func(assert *assert.Assertions) {
		test.DefaultVerify(assert)

		projectID := test.GetStringOutput("project_id")

		// Run gcloud command to verify the binding
		op := gcloud.Run(t, fmt.Sprintf("projects get-iam-policy %s", projectID))

		// Search for the specific role in the policy array
		found := false
		for _, binding := range op.Get("bindings").Array() {
			if binding.Get("role").String() == "roles/viewer" {
				found = true
				break
			}
		}
		assert.True(found, "should have roles/viewer binding")
	})

	test.Test()
}
