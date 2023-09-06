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

package appengine_test

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/lorioux/google-beta/google-beta/acctest"
	"github.com/lorioux/google-beta/google-beta/envvar"
)

func TestAccAppEngineFlexibleAppVersion_appEngineFlexibleAppVersionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAppEngineFlexibleAppVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineFlexibleAppVersion_appEngineFlexibleAppVersionExample(context),
			},
			{
				ResourceName:            "google_app_engine_flexible_app_version.myapp_v1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"beta_settings", "env_variables", "entrypoint", "service", "noop_on_destroy", "deployment.0.zip"},
			},
		},
	})
}

func testAccAppEngineFlexibleAppVersion_appEngineFlexibleAppVersionExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "my_project" {
  name = "tf-test-appeng-flex%{random_suffix}"
  project_id = "tf-test-appeng-flex%{random_suffix}"
  org_id = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_app_engine_application" "app" {
  project     = google_project.my_project.project_id
  location_id = "us-central"
}

resource "google_project_service" "service" {
  project = google_project.my_project.project_id
  service = "appengineflex.googleapis.com"

  disable_dependent_services = false
}

resource "google_service_account" "custom_service_account" {
  project      = google_project_service.service.project
  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Custom Service Account"
}

resource "google_project_iam_member" "gae_api" {
  project = google_project_service.service.project
  role    = "roles/compute.networkUser"
  member  = "serviceAccount:${google_service_account.custom_service_account.email}"
}

resource "google_project_iam_member" "logs_writer" {
  project = google_project_service.service.project
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.custom_service_account.email}"
}

resource "google_project_iam_member" "storage_viewer" {
  project = google_project_service.service.project
  role    = "roles/storage.objectViewer"
  member  = "serviceAccount:${google_service_account.custom_service_account.email}"
}

resource "google_app_engine_flexible_app_version" "myapp_v1" {
  version_id = "v1"
  project    = google_project_iam_member.gae_api.project
  service    = "default"
  runtime    = "nodejs"

  entrypoint {
    shell = "node ./app.js"
  }

  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }
  }

  liveness_check {
    path = "/"
  }

  readiness_check {
    path = "/"
  }

  env_variables = {
    port = "8080"
  }

  handlers {
    url_regex        = ".*\\/my-path\\/*"
    security_level   = "SECURE_ALWAYS"
    login            = "LOGIN_REQUIRED"
    auth_fail_action = "AUTH_FAIL_ACTION_REDIRECT"

    static_files {
      path = "my-other-path"
      upload_path_regex = ".*\\/my-path\\/*"
    }
  }

  automatic_scaling {
    cool_down_period = "120s"
    cpu_utilization {
      target_utilization = 0.5
    }
  }

  noop_on_destroy = true
  service_account = google_service_account.custom_service_account.email
}

resource "google_storage_bucket" "bucket" {
  project  = google_project.my_project.project_id
  name     = "tf-test-appengine-static-content%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/hello-world.zip"
}
`, context)
}

func testAccCheckAppEngineFlexibleAppVersionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_app_engine_flexible_app_version" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}
