// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package iambeta_test

import (
	"fmt"
	"testing"

	"github.com/lorioux/google-beta/google-beta/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIAMBetaWorkloadIdentityPool_full(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(t, 10)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPool_full(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPool_update(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIAMBetaWorkloadIdentityPool_minimal(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(t, 10)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPool_minimal(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPool_update(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPool_full(suffix string) string {
	return fmt.Sprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%s"
  display_name              = "Name of pool"
  description               = "Identity pool for automated test"
  disabled                  = true
}
`, suffix)
}

func testAccIAMBetaWorkloadIdentityPool_minimal(suffix string) string {
	return fmt.Sprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%s"
}
`, suffix)
}

func testAccIAMBetaWorkloadIdentityPool_update(suffix string) string {
	return fmt.Sprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%s"
  display_name              = "Updated name of pool"
  description               = "Updated description"
  disabled                  = false
}
`, suffix)
}
