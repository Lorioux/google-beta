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

package networkservices_test

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

func TestAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_service.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "dest" {
  name          = "tf-test-my-bucket%{random_suffix}"
  location      = "US"
  force_destroy = true
}

resource "google_network_services_edge_cache_origin" "instance" {
  name                 = "tf-test-my-origin%{random_suffix}"
  origin_address       = google_storage_bucket.dest.url
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  timeout {
    connect_timeout = "10s"
  }
}

resource "google_network_services_edge_cache_service" "instance" {
  name                 = "tf-test-my-service%{random_suffix}"
  description          = "some description"
  routing {
    host_rule {
      description = "host rule description"
      hosts = ["sslcert.tf-test.club"]
      path_matcher = "routes"
    }
    path_matcher {
      name = "routes"
      route_rule {
        description = "a route rule to match against"
        priority = 1
        match_rule {
          prefix_match = "/"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
              cache_mode = "CACHE_ALL_STATIC"
              default_ttl = "3600s"
          }
        }
        header_action {
          response_header_to_add {
            header_name = "x-cache-status"
            header_value = "{cdn_cache_status}"
          }
        }
      }
    }
  }
}
`, context)
}

func TestAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_service.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "dest" {
  name          = "tf-test-my-bucket%{random_suffix}"
  location      = "US"
  force_destroy = true
}

resource "google_network_services_edge_cache_origin" "google" {
  name                 = "tf-test-origin-google%{random_suffix}"
  origin_address       = "google.com"
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  timeout {
    connect_timeout = "10s"
  }
}

resource "google_network_services_edge_cache_origin" "instance" {
  name                 = "tf-test-my-origin%{random_suffix}"
  origin_address       = google_storage_bucket.dest.url
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  timeout {
    connect_timeout = "10s"
  }
}

resource "google_network_services_edge_cache_service" "instance" {
  name                 = "tf-test-my-service%{random_suffix}"
  description          = "some description"
  disable_quic         = true
  disable_http2        = true
  labels = {
    a = "b"
  }

  routing {
    host_rule {
      description = "host rule description"
      hosts = ["sslcert.tf-test.club"]
      path_matcher = "routes"
    }
    host_rule {
      description = "host rule2"
      hosts = ["sslcert.tf-test2.club"]
      path_matcher = "routes"
    }

    host_rule {
      description = "host rule3"
      hosts = ["sslcert.tf-test3.club"]
      path_matcher = "routesAdvanced"
    }

    path_matcher {
      name = "routes"
      route_rule {
        description = "a route rule to match against"
        priority = 1
        match_rule {
          prefix_match = "/"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
              cache_mode = "CACHE_ALL_STATIC"
              default_ttl = "3600s"
          }
        }
        header_action {
          response_header_to_add {
            header_name = "x-cache-status"
            header_value = "{cdn_cache_status}"
          }
        }
      }
    }

    path_matcher {
      name = "routesAdvanced"
      description = "an advanced ruleset"
      route_rule {
        description = "an advanced route rule to match against"
        priority = 1
        match_rule {
          prefix_match = "/potato/"
          query_parameter_match {
            name = "debug"
            present_match = true
          }
          query_parameter_match {
            name = "state"
            exact_match = "debug"
          }
        }
        match_rule {
          full_path_match = "/apple"
        }
        header_action {
          request_header_to_add {
            header_name = "debug"
            header_value = "true"
            replace = true
          }
          request_header_to_add {
            header_name = "potato"
            header_value = "plant"
          }
          response_header_to_add {
            header_name = "potato"
            header_value = "plant"
            replace = true
          }
          request_header_to_remove {
            header_name = "prod"
          }
          response_header_to_remove {
            header_name = "prod"
          }
        }

        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
              cache_mode = "CACHE_ALL_STATIC"
              default_ttl = "3800s"
              client_ttl = "3600s"
              max_ttl = "9000s"
              cache_key_policy {
                include_protocol = true
                exclude_host = true
                included_query_parameters = ["apple", "dev", "santa", "claus"]
                included_header_names = ["banana"]
                included_cookie_names = ["orange"]
              }
              negative_caching = true
              signed_request_mode = "DISABLED"
              negative_caching_policy = {
                "500" = "3000s"
              }
          }
          url_rewrite {
            path_prefix_rewrite = "/dev"
            host_rewrite = "dev.club"
          }
          cors_policy {
            max_age = "2500s"
            allow_credentials = true
            allow_origins = ["*"]
            allow_methods = ["GET"]
            allow_headers = ["dev"]
            expose_headers = ["prod"]
          }
        }
      }
      route_rule {
        description = "a second route rule to match against"
        priority = 2
        match_rule {
          full_path_match = "/yay"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
            cache_mode = "CACHE_ALL_STATIC"
            default_ttl = "3600s"
            cache_key_policy {
              excluded_query_parameters = ["dev"]
            }
          }
          cors_policy {
            max_age = "3000s"
            allow_headers = ["dev"]
            disabled = true
          }
        }
      }
    }
  }

  log_config {
    enable = true
    sample_rate = 0.01
  }
}
`, context)
}

func TestAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceDualTokenExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceDualTokenExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_service.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheService_networkServicesEdgeCacheServiceDualTokenExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-name%{random_suffix}"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "secret-data"
}

resource "google_network_services_edge_cache_keyset" "keyset" {
  name        = "tf-test-keyset-name%{random_suffix}"
  description = "The default keyset"
  public_key {
    id      = "my-public-key"
    managed = true
  }
  validation_shared_keys {
    secret_version = google_secret_manager_secret_version.secret-version-basic.id
  }
}

resource "google_network_services_edge_cache_origin" "instance" {
  name                 = "tf-test-my-origin%{random_suffix}"
  origin_address       = "gs://media-edge-default"
  description          = "The default bucket for media edge test"
}

resource "google_network_services_edge_cache_service" "instance" {
  name                 = "tf-test-my-service%{random_suffix}"
  description          = "some description"
  routing {
    host_rule {
      description = "host rule description"
      hosts = ["sslcert.tf-test.club"]
      path_matcher = "routes"
    }
    path_matcher {
      name = "routes"
      route_rule {
        description = "a route rule to match against master playlist"
        priority = 1
        match_rule {
          path_template_match = "/master.m3u8"
	}	
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
	    signed_request_mode = "REQUIRE_TOKENS"
	    signed_request_keyset = google_network_services_edge_cache_keyset.keyset.id
	    signed_token_options {
	      token_query_parameter = "edge-cache-token"
	    }
	    signed_request_maximum_expiration_ttl = "600s"
	    add_signatures {
	      actions = ["GENERATE_COOKIE"]
	      keyset = google_network_services_edge_cache_keyset.keyset.id
	      copied_parameters = ["PathGlobs", "SessionID"]
	    }
          }
        }
      }
      route_rule {
        description = "a route rule to match against all playlists"
        priority = 2
        match_rule {
          path_template_match = "/*.m3u8"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
	    signed_request_mode = "REQUIRE_TOKENS"
	    signed_request_keyset = google_network_services_edge_cache_keyset.keyset.id
	    signed_token_options {
	      token_query_parameter = "hdnts"
	      allowed_signature_algorithms = ["ED25519", "HMAC_SHA_256", "HMAC_SHA1"]
	    }
	    add_signatures {
	      actions = ["GENERATE_TOKEN_HLS_COOKIELESS"]
	      keyset = google_network_services_edge_cache_keyset.keyset.id
	      token_ttl = "1200s"
	      token_query_parameter = "hdntl"
	      copied_parameters = ["URLPrefix"]
	    }
          }
        }
      }
      route_rule {
        description = "a route rule to match against"
        priority = 3
        match_rule {
          path_template_match = "/**.m3u8"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
	    signed_request_mode = "REQUIRE_TOKENS"
	    signed_request_keyset = google_network_services_edge_cache_keyset.keyset.id
	    signed_token_options {
	      token_query_parameter = "hdntl"
	    }
	    add_signatures {
	      actions = ["PROPAGATE_TOKEN_HLS_COOKIELESS"]
	      token_query_parameter = "hdntl"
	    }
          }
        }
      }
    }
  }
}
`, context)
}

func testAccCheckNetworkServicesEdgeCacheServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_edge_cache_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheServices/{{name}}")
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
				return fmt.Errorf("NetworkServicesEdgeCacheService still exists at %s", url)
			}
		}

		return nil
	}
}
