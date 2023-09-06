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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/lorioux/google-beta/google-beta/acctest"
	"github.com/lorioux/google-beta/google-beta/envvar"
)

func TestAccComputeBackendBucketIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucketIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendBuckets/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-image-backend-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeBackendBucketIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendBuckets/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-image-backend-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendBucketIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeBackendBucketIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendBuckets/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-image-backend-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendBucketIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucketIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_compute_backend_bucket_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_compute_backend_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendBuckets/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-image-backend-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendBucketIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendBuckets/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-image-backend-bucket%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucketIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}

resource "google_compute_backend_bucket_iam_member" "foo" {
  project = google_compute_backend_bucket.image_backend.project
  name = google_compute_backend_bucket.image_backend.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeBackendBucketIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_backend_bucket_iam_policy" "foo" {
  project = google_compute_backend_bucket.image_backend.project
  name = google_compute_backend_bucket.image_backend.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_compute_backend_bucket_iam_policy" "foo" {
  project = google_compute_backend_bucket.image_backend.project
  name = google_compute_backend_bucket.image_backend.name
  depends_on = [
    google_compute_backend_bucket_iam_policy.foo
  ]
}
`, context)
}

func testAccComputeBackendBucketIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}

data "google_iam_policy" "foo" {
}

resource "google_compute_backend_bucket_iam_policy" "foo" {
  project = google_compute_backend_bucket.image_backend.project
  name = google_compute_backend_bucket.image_backend.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeBackendBucketIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}

resource "google_compute_backend_bucket_iam_binding" "foo" {
  project = google_compute_backend_bucket.image_backend.project
  name = google_compute_backend_bucket.image_backend.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeBackendBucketIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}

resource "google_compute_backend_bucket_iam_binding" "foo" {
  project = google_compute_backend_bucket.image_backend.project
  name = google_compute_backend_bucket.image_backend.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
