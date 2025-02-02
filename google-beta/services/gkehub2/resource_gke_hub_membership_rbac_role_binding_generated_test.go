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

package gkehub2_test

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

func TestAccGKEHub2MembershipRBACRoleBinding_gkehubMembershipRbacRoleBindingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGKEHub2MembershipRBACRoleBindingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHub2MembershipRBACRoleBinding_gkehubMembershipRbacRoleBindingBasicExample(context),
			},
			{
				ResourceName:            "google_gke_hub_membership_rbac_role_binding.membershiprbacrolebinding",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"membership_rbac_role_binding_id", "membership_id", "location"},
			},
		},
	})
}

func testAccGKEHub2MembershipRBACRoleBinding_gkehubMembershipRbacRoleBindingBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  provider = google-beta
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membershiprbacrolebinding" {
  provider = google-beta
  membership_id = "tf-test-membership%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }

  depends_on = [google_container_cluster.primary]
}

resource "google_gke_hub_membership_rbac_role_binding" "membershiprbacrolebinding" {
  provider = google-beta
  membership_rbac_role_binding_id = "tf-test-membership-rbac-role-binding%{random_suffix}"
  membership_id = "tf-test-membership%{random_suffix}"
  user = "service-${data.google_project.project.number}@gcp-sa-anthossupport.iam.gserviceaccount.com"
  role {
    predefined_role = "ANTHOS_SUPPORT"
  }
  location = "global"
  depends_on = [google_gke_hub_membership.membershiprbacrolebinding]
}

data "google_project" "project" {
  provider = google-beta
}
`, context)
}

func testAccCheckGKEHub2MembershipRBACRoleBindingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gke_hub_membership_rbac_role_binding" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}")
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
				return fmt.Errorf("GKEHub2MembershipRBACRoleBinding still exists at %s", url)
			}
		}

		return nil
	}
}
