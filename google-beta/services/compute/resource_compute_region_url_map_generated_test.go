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

	"github.com/lorioux/google-beta/google-beta/acctest"
	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func TestAccComputeRegionUrlMap_regionUrlMapBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_url_map.regionurlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service", "region"},
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  region = "us-central1"

  name        = "regionurlmap%{random_suffix}"
  description = "a description"

  default_service = google_compute_region_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      service = google_compute_region_backend_service.home.id
    }

    path_rule {
      paths   = ["/login"]
      service = google_compute_region_backend_service.login.id
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "login" {
  region = "us-central1"

  name        = "login%{random_suffix}"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_backend_service" "home" {
  region = "us-central1"

  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"

  name               = "tf-test-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  http_health_check {
    port         = 80
    request_path = "/"
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapDefaultRouteActionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapDefaultRouteActionExample(context),
			},
			{
				ResourceName:            "google_compute_region_url_map.regionurlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service", "region"},
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapDefaultRouteActionExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  region = "us-central1"

  name        = "regionurlmap%{random_suffix}"
  description = "a description"

  default_route_action {
    retry_policy {
      retry_conditions = [
        "5xx",
        "gateway-error",
      ]
      num_retries = 3
      per_try_timeout {
        seconds = 0
        nanos = 500
      }
    }
    request_mirror_policy {
      backend_service = google_compute_region_backend_service.home.id
    }
    weighted_backend_services {
      backend_service = google_compute_region_backend_service.login.id
      weight = 200
      header_action {
        request_headers_to_add {
          header_name = "foo-request-1"
          header_value = "bar"
          replace = true
        }
        request_headers_to_remove = ["fizz"]
        response_headers_to_add {
          header_name = "foo-response-1"
          header_value = "bar"
          replace = true
        }
        response_headers_to_remove = ["buzz"]
      }
    }
    weighted_backend_services {
      backend_service = google_compute_region_backend_service.home.id
      weight = 100
      header_action {
        request_headers_to_add {
          header_name = "foo-request-1"
          header_value = "bar"
          replace = true
        }
        request_headers_to_add {
          header_name = "foo-request-2"
          header_value = "bar"
          replace = true
        }
        request_headers_to_remove = ["fizz"]
        response_headers_to_add {
          header_name = "foo-response-2"
          header_value = "bar"
          replace = true
        }
        response_headers_to_add {
          header_name = "foo-response-1"
          header_value = "bar"
          replace = true
        }
        response_headers_to_remove = ["buzz"]
      }
    }
    url_rewrite {
      host_rewrite = "dev.example.com"
      path_prefix_rewrite = "/v1/api/"
    }
  
    cors_policy {
      disabled = false
      allow_credentials = true
      allow_headers = [
        "foobar"
      ]
      allow_methods = [
        "GET",
        "POST",
      ]
      allow_origins = [
        "example.com"
      ]
      expose_headers = [
        "foobar"
      ]
      max_age = 60
    }
    fault_injection_policy {
      delay {
        fixed_delay {
          seconds = 0
          nanos = 500
        }
        percentage = 0.5
      }
      abort {
        http_status = 500
        percentage = 0.5
      }
    }
    timeout {
      seconds = 0
      nanos = 500
    }
  }

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      service = google_compute_region_backend_service.home.id
    }

    path_rule {
      paths   = ["/login"]
      service = google_compute_region_backend_service.login.id
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "login" {
  region = "us-central1"

  name        = "login%{random_suffix}"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_backend_service" "home" {
  region = "us-central1"

  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"

  name               = "tf-test-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  http_health_check {
    port         = 80
    request_path = "/"
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbPathExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbPathExample(context),
			},
			{
				ResourceName:            "google_compute_region_url_map.regionurlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service", "region"},
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbPathExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  name        = "regionurlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_region_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      route_action {
        cors_policy {
          allow_credentials = true
          allow_headers = ["Allowed content"]
          allow_methods = ["GET"]
          allow_origins = ["Allowed origin"]
          expose_headers = ["Exposed header"]
          max_age = 30
          disabled = false
        }
        fault_injection_policy {
          abort {
            http_status = 234
            percentage = 5.6
          }
          delay {
            fixed_delay {
              seconds = 0
              nanos = 50000
            }
            percentage = 7.8
          }
        }
        request_mirror_policy {
          backend_service = google_compute_region_backend_service.home.id
        }
        retry_policy {
          num_retries = 4
          per_try_timeout {
            seconds = 30
          }
          retry_conditions = ["5xx", "deadline-exceeded"]
        }
        timeout {
          seconds = 20
          nanos = 750000000
        }
        url_rewrite {
          host_rewrite = "dev.example.com"
          path_prefix_rewrite = "/v1/api/"
        }
        weighted_backend_services {
          backend_service = google_compute_region_backend_service.home.id
          weight = 400
          header_action {
            request_headers_to_remove = ["RemoveMe"]
            request_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = true
            }
            response_headers_to_remove = ["RemoveMe"]
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = false
            }
          }
        }
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbPathPartialExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbPathPartialExample(context),
			},
			{
				ResourceName:            "google_compute_region_url_map.regionurlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service", "region"},
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbPathPartialExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  name        = "regionurlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_region_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      route_action {
        retry_policy {
          num_retries = 4
          per_try_timeout {
            seconds = 30
          }
          retry_conditions = ["5xx", "deadline-exceeded"]
        }
        timeout {
          seconds = 20
          nanos = 750000000
        }
        url_rewrite {
          host_rewrite = "dev.example.com"
          path_prefix_rewrite = "/v1/api/"
        }
        weighted_backend_services {
          backend_service = google_compute_region_backend_service.home.id
          weight = 400
          header_action {
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = false
            }
          }
        }
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbRouteExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbRouteExample(context),
			},
			{
				ResourceName:            "google_compute_region_url_map.regionurlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service", "region"},
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbRouteExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  name            = "regionurlmap%{random_suffix}"
  description     = "a description"
  default_service = google_compute_region_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    route_rules {
      priority = 1
      header_action {
        request_headers_to_remove = ["RemoveMe2"]
        request_headers_to_add {
          header_name = "AddSomethingElse"
          header_value = "MyOtherValue"
          replace = true
        }
        response_headers_to_remove = ["RemoveMe3"]
        response_headers_to_add {
          header_name = "AddMe"
          header_value = "MyValue"
          replace = false
        }
      }
      match_rules {
        full_path_match = "a full path"
        header_matches {
          header_name = "someheader"
          exact_match = "match this exactly"
          invert_match = true
        }
        ignore_case = true
        metadata_filters {
          filter_match_criteria = "MATCH_ANY"
          filter_labels {
            name = "PLANET"
            value = "MARS"
          }
        }
        query_parameter_matches {
          name = "a query parameter"
          present_match = true
        }
      }
      url_redirect {
        host_redirect = "A host"
        https_redirect = false
        path_redirect = "some/path"
        redirect_response_code = "TEMPORARY_REDIRECT"
        strip_query = true
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name     = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbRoutePartialExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbRoutePartialExample(context),
			},
			{
				ResourceName:            "google_compute_region_url_map.regionurlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service", "region"},
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbRoutePartialExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  name        = "regionurlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_region_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    route_rules {
      priority = 1
      service = google_compute_region_backend_service.home.id
      header_action {
        request_headers_to_remove = ["RemoveMe2"]
      }
      match_rules {
        full_path_match = "a full path"
        header_matches {
          header_name = "someheader"
          exact_match = "match this exactly"
          invert_match = true
        }
        query_parameter_matches {
          name = "a query parameter"
          present_match = true
        }
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func testAccCheckComputeRegionUrlMapDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_url_map" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/urlMaps/{{name}}")
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
				return fmt.Errorf("ComputeRegionUrlMap still exists at %s", url)
			}
		}

		return nil
	}
}
