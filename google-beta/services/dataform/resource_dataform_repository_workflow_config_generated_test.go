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

package dataform_test

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

func TestAccDataformRepositoryWorkflowConfig_dataformRepositoryWorkflowConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDataformRepositoryWorkflowConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataformRepositoryWorkflowConfig_dataformRepositoryWorkflowConfigExample(context),
			},
			{
				ResourceName:            "google_dataform_repository_workflow_config.workflow",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "repository"},
			},
		},
	})
}

func testAccDataformRepositoryWorkflowConfig_dataformRepositoryWorkflowConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name     = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider  = google-beta
  secret_id = "tf_test_my_secret%{random_suffix}"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret   = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "repository" {
  provider = google-beta
  name     = "tf_test_dataform_repository%{random_suffix}"
  region   = "us-central1"

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}

resource "google_dataform_repository_release_config" "release_config" {
  provider = google-beta

  project    = google_dataform_repository.repository.project
  region     = google_dataform_repository.repository.region
  repository = google_dataform_repository.repository.name

  name          = "tf_test_my_release%{random_suffix}"
  git_commitish = "main"
  cron_schedule = "0 7 * * *"
  time_zone     = "America/New_York"

  code_compilation_config {
    default_database = "gcp-example-project"
    default_schema   = "example-dataset"
    default_location = "us-central1"
    assertion_schema = "example-assertion-dataset"
    database_suffix  = ""
    schema_suffix    = ""
    table_prefix     = ""
    vars = {
      var1 = "value"
    }
  }
}

resource "google_service_account" "dataform_sa" {
  provider     = google-beta
  account_id   = "dataform-workflow-sa"
  display_name = "Dataform Service Account"
}

resource "google_dataform_repository_workflow_config" "workflow" {
  provider = google-beta

  project        = google_dataform_repository.repository.project
  region         = google_dataform_repository.repository.region
  repository     = google_dataform_repository.repository.name
  name           = "tf_test_my_workflow%{random_suffix}"
  release_config = google_dataform_repository_release_config.release_config.id

  invocation_config {
    included_targets {
      database = "gcp-example-project"
      schema   = "example-dataset"
      name     = "target_1"
    }
    included_targets {
      database = "gcp-example-project"
      schema   = "example-dataset"
      name     = "target_2"
    }
    included_tags                            = ["tag_1"]
    transitive_dependencies_included         = true
    transitive_dependents_included           = true
    fully_refresh_incremental_tables_enabled = false
    service_account                          = google_service_account.dataform_sa.email
  }

  cron_schedule   = "0 7 * * *"
  time_zone       = "America/New_York"
}
`, context)
}

func testAccCheckDataformRepositoryWorkflowConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataform_repository_workflow_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
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
				return fmt.Errorf("DataformRepositoryWorkflowConfig still exists at %s", url)
			}
		}

		return nil
	}
}
