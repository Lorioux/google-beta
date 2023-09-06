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

func TestAccGKEHub2ScopeRBACRoleBinding_gkehubScopeRbacRoleBindingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEHub2ScopeRBACRoleBindingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHub2ScopeRBACRoleBinding_gkehubScopeRbacRoleBindingBasicExample(context),
			},
			{
				ResourceName:            "google_gke_hub_scope_rbac_role_binding.scoperbacrolebinding",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"scope_rbac_role_binding_id", "scope_id"},
			},
		},
	})
}

func testAccGKEHub2ScopeRBACRoleBinding_gkehubScopeRbacRoleBindingBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gke_hub_scope" "scoperbacrolebinding" {
  scope_id = "tf-test-scope%{random_suffix}"
}

resource "google_gke_hub_scope_rbac_role_binding" "scoperbacrolebinding" {
  scope_rbac_role_binding_id = "tf-test-scope-rbac-role-binding%{random_suffix}"
  scope_id = "tf-test-scope%{random_suffix}"
  user = "test-email@gmail.com"
  role {
    predefined_role = "ADMIN"
  }
  labels = {
      key = "value" 
  }
  depends_on = [google_gke_hub_scope.scoperbacrolebinding]
}
`, context)
}

func testAccCheckGKEHub2ScopeRBACRoleBindingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gke_hub_scope_rbac_role_binding" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
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
				return fmt.Errorf("GKEHub2ScopeRBACRoleBinding still exists at %s", url)
			}
		}

		return nil
	}
}
