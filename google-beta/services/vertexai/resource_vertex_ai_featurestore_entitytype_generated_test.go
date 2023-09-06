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

package vertexai_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/lorioux/google-beta/google-beta/acctest"
	"github.com/lorioux/google-beta/google-beta/envvar"
	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func TestAccVertexAIFeaturestoreEntitytype_vertexAiFeaturestoreEntitytypeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"kms_key_name":    acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIFeaturestoreEntitytypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreEntitytype_vertexAiFeaturestoreEntitytypeExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_featurestore_entitytype.entity",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "etag", "featurestore"},
			},
		},
	})
}

func testAccVertexAIFeaturestoreEntitytype_vertexAiFeaturestoreEntitytypeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  description = "test description"
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval_days = 1
      staleness_days = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}
`, context)
}

func TestAccVertexAIFeaturestoreEntitytype_vertexAiFeaturestoreEntitytypeWithBetaFieldsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"kms_key_name":    acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVertexAIFeaturestoreEntitytypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreEntitytype_vertexAiFeaturestoreEntitytypeWithBetaFieldsExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_featurestore_entitytype.entity",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "etag", "featurestore"},
			},
		},
	})
}

func testAccVertexAIFeaturestoreEntitytype_vertexAiFeaturestoreEntitytypeWithBetaFieldsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  provider = google-beta
  name     = "terraform2%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
}

resource "google_vertex_ai_featurestore_entitytype" "entity" {
  provider = google-beta
  name     = "terraform2%{random_suffix}"
  labels = {
    foo = "bar"
  }
  featurestore = google_vertex_ai_featurestore.featurestore.id
  monitoring_config {
    snapshot_analysis {
      disabled = false
      monitoring_interval = "86400s"
    }

    categorical_threshold_config {
      value = 0.3
    }

    numerical_threshold_config {
      value = 0.3
    }
  }
  offline_storage_ttl_days = 30
}
`, context)
}

func testAccCheckVertexAIFeaturestoreEntitytypeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_featurestore_entitytype" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
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
				return fmt.Errorf("VertexAIFeaturestoreEntitytype still exists at %s", url)
			}
		}

		return nil
	}
}
