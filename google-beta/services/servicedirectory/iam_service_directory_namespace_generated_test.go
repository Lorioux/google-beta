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

package servicedirectory_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/lorioux/google-beta/google-beta/acctest"
)

func TestAccServiceDirectoryNamespaceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryNamespaceIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccServiceDirectoryNamespaceIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccServiceDirectoryNamespaceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccServiceDirectoryNamespaceIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccServiceDirectoryNamespaceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryNamespaceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_service_directory_namespace_iam_policy.foo", "policy_data"),
			},
			{
				Config: testAccServiceDirectoryNamespaceIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccServiceDirectoryNamespaceIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

resource "google_service_directory_namespace_iam_member" "foo" {
  provider = google-beta
  name = google_service_directory_namespace.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccServiceDirectoryNamespaceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_service_directory_namespace_iam_policy" "foo" {
  provider = google-beta
  name = google_service_directory_namespace.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_service_directory_namespace_iam_policy" "foo" {
  provider = google-beta
  name = google_service_directory_namespace.example.name
  depends_on = [
    google_service_directory_namespace_iam_policy.foo
  ]
}
`, context)
}

func testAccServiceDirectoryNamespaceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_service_directory_namespace_iam_policy" "foo" {
  provider = google-beta
  name = google_service_directory_namespace.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccServiceDirectoryNamespaceIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

resource "google_service_directory_namespace_iam_binding" "foo" {
  provider = google-beta
  name = google_service_directory_namespace.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccServiceDirectoryNamespaceIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

resource "google_service_directory_namespace_iam_binding" "foo" {
  provider = google-beta
  name = google_service_directory_namespace.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
