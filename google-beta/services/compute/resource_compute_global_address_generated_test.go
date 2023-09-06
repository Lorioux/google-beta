// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeGlobalAddress_globalAddressBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalAddress_globalAddressBasicExample(context),
			},
			{
				ResourceName:            "google_compute_global_address.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network"},
			},
		},
	})
}

func testAccComputeGlobalAddress_globalAddressBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_global_address" "default" {
  name = "tf-test-global-appserver-ip%{random_suffix}"
}
`, context)
}

func TestAccComputeGlobalAddress_globalAddressPrivateServicesConnectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalAddress_globalAddressPrivateServicesConnectExample(context),
			},
			{
				ResourceName:            "google_compute_global_address.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network"},
			},
		},
	})
}

func testAccComputeGlobalAddress_globalAddressPrivateServicesConnectExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_global_address" "default" {
  provider      = google-beta
  name          = "tf-test-global-psconnect-ip%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.network.id
  address       = "100.100.100.105"
}

resource "google_compute_network" "network" {
  provider      = google-beta
  name          = "tf-test-my-network-name%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckComputeGlobalAddressDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_global_address" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/addresses/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeGlobalAddress still exists at %s", url)
			}
		}

		return nil
	}
}
