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

package healthcare_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/lorioux/google-beta/google-beta/acctest"
	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func TestAccHealthcareConsentStore_healthcareConsentStoreBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareConsentStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStore_healthcareConsentStoreBasicExample(context),
			},
			{
				ResourceName:            "google_healthcare_consent_store.my-consent",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "dataset"},
			},
		},
	})
}

func testAccHealthcareConsentStore_healthcareConsentStoreBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}
`, context)
}

func TestAccHealthcareConsentStore_healthcareConsentStoreFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareConsentStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStore_healthcareConsentStoreFullExample(context),
			},
			{
				ResourceName:            "google_healthcare_consent_store.my-consent",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "dataset"},
			},
		},
	})
}

func testAccHealthcareConsentStore_healthcareConsentStoreFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"

  enable_consent_create_on_update = true
  default_consent_ttl             = "90000s"

  labels = {
    "label1" = "labelvalue1"
  }
}
`, context)
}

func TestAccHealthcareConsentStore_healthcareConsentStoreIamExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareConsentStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStore_healthcareConsentStoreIamExample(context),
			},
			{
				ResourceName:            "google_healthcare_consent_store.my-consent",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "dataset"},
			},
		},
	})
}

func testAccHealthcareConsentStore_healthcareConsentStoreIamExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_service_account" "test-account" {
  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_healthcare_consent_store_iam_member" "test-iam" {
  dataset          = google_healthcare_dataset.dataset.id
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role             = "roles/editor"
  member           = "serviceAccount:${google_service_account.test-account.email}"
}
`, context)
}

func testAccCheckHealthcareConsentStoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_consent_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{HealthcareBasePath}}{{dataset}}/consentStores/{{name}}")
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
				return fmt.Errorf("HealthcareConsentStore still exists at %s", url)
			}
		}

		return nil
	}
}
