package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	computeBeta "google.golang.org/api/compute/v0.beta"
)

func TestAccComputeGlobalForwardingRule_basic(t *testing.T) {
	t.Parallel()

	fr := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy1 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy2 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	backend := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	hc := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeGlobalForwardingRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeGlobalForwardingRule_basic1(fr, proxy1, proxy2, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeGlobalForwardingRuleExists(
						"google_compute_global_forwarding_rule.foobar"),

					// Unlike GlobalAddress, IpVersion is "" instead of "IPV4" when this is made with a v1 API.
					testAccCheckComputeBetaGlobalForwardingRuleIpVersion("google_compute_global_forwarding_rule.foobar", ""),
				),
			},
			resource.TestStep{
				ResourceName:      "google_compute_global_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeGlobalForwardingRule_update(t *testing.T) {
	t.Parallel()

	fr := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy1 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy2 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	backend := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	hc := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeGlobalForwardingRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeGlobalForwardingRule_basic1(fr, proxy1, proxy2, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeGlobalForwardingRuleExists(
						"google_compute_global_forwarding_rule.foobar"),
				),
			},

			resource.TestStep{
				Config: testAccComputeGlobalForwardingRule_basic2(fr, proxy1, proxy2, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeGlobalForwardingRuleExists(
						"google_compute_global_forwarding_rule.foobar"),
				),
			},
		},
	})
}

func TestAccComputeGlobalForwardingRule_ipv6(t *testing.T) {
	t.Parallel()

	var frule computeBeta.ForwardingRule

	fr := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy1 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy2 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	backend := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	hc := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeGlobalForwardingRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeGlobalForwardingRule_ipv6(fr, proxy1, proxy2, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeBetaGlobalForwardingRuleExists(
						"google_compute_global_forwarding_rule.foobar", &frule),
					testAccCheckComputeBetaGlobalForwardingRuleIpVersion("google_compute_global_forwarding_rule.foobar", "IPV6"),
				),
			},
			resource.TestStep{
				ResourceName:      "google_compute_global_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeGlobalForwardingRule_labels(t *testing.T) {
	t.Parallel()

	var frule computeBeta.ForwardingRule

	fr := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy1 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	proxy2 := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	backend := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	hc := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeGlobalForwardingRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeGlobalForwardingRule_labels(fr, proxy1, proxy2, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeBetaGlobalForwardingRuleExists(
						"google_compute_global_forwarding_rule.foobar", &frule),
					testAccCheckComputeBetaGlobalForwardingRuleHasLabel(&frule, "my-label", "my-label-value"),
					testAccCheckComputeBetaGlobalForwardingRuleHasLabel(&frule, "my-second-label", "my-second-label-value"),
					testAccCheckComputeBetaGlobalForwardingRuleHasCorrectLabelFingerprint(&frule, "google_compute_global_forwarding_rule.foobar"),
				),
			},
			resource.TestStep{
				Config: testAccComputeGlobalForwardingRule_labelsUpdated(fr, proxy1, proxy2, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeBetaGlobalForwardingRuleExists(
						"google_compute_global_forwarding_rule.foobar", &frule),
					testAccCheckComputeBetaGlobalForwardingRuleHasLabel(&frule, "my-label", "my-label-value"),
					testAccCheckComputeBetaGlobalForwardingRuleHasLabel(&frule, "my-third-label", "my-third-label-value"),
					testAccCheckComputeBetaGlobalForwardingRuleHasCorrectLabelFingerprint(&frule, "google_compute_global_forwarding_rule.foobar"),
				),
			},
		},
	})
}

func testAccCheckComputeGlobalForwardingRuleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_global_forwarding_rule" {
			continue
		}

		_, err := config.clientCompute.GlobalForwardingRules.Get(
			config.Project, rs.Primary.ID).Do()
		if err == nil {
			return fmt.Errorf("Global Forwarding Rule still exists")
		}
	}

	return nil
}

func testAccCheckComputeGlobalForwardingRuleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)

		found, err := config.clientCompute.GlobalForwardingRules.Get(
			config.Project, rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.ID {
			return fmt.Errorf("Global Forwarding Rule not found")
		}

		return nil
	}
}

func testAccCheckComputeBetaGlobalForwardingRuleExists(n string, frule *computeBeta.ForwardingRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)

		found, err := config.clientComputeBeta.GlobalForwardingRules.Get(
			config.Project, rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.ID {
			return fmt.Errorf("Global Forwarding Rule not found")
		}

		*frule = *found

		return nil
	}
}

func testAccCheckComputeBetaGlobalForwardingRuleIpVersion(n, version string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)

		frule, err := config.clientComputeBeta.GlobalForwardingRules.Get(config.Project, rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		if frule.IpVersion != version {
			return fmt.Errorf("Expected IP version to be %s, got %s", version, frule.IpVersion)
		}

		return nil
	}
}
func testAccCheckComputeBetaGlobalForwardingRuleHasLabel(frule *computeBeta.ForwardingRule, key, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		val, ok := frule.Labels[key]
		if !ok {
			return fmt.Errorf("label with key %s not found", key)
		}

		if val != value {
			return fmt.Errorf("label value did not match for key %s: expected %s but found %s", key, value, val)
		}
		return nil
	}
}

func testAccCheckComputeBetaGlobalForwardingRuleHasCorrectLabelFingerprint(
	frule *computeBeta.ForwardingRule, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tfLabelFingerprint := s.RootModule().Resources[resourceName].Primary.Attributes["label_fingerprint"]
		remoteLabelFingerprint := frule.LabelFingerprint

		if tfLabelFingerprint != remoteLabelFingerprint {
			return fmt.Errorf("Label fingerprint mismatch: remote has %#v but terraform has %#v",
				remoteLabelFingerprint, tfLabelFingerprint)
		}
		return nil
	}
}

func testAccComputeGlobalForwardingRule_basic1(fr, proxy1, proxy2, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
	resource "google_compute_global_forwarding_rule" "foobar" {
		description = "Resource created for Terraform acceptance testing"
		ip_protocol = "TCP"
		name = "%s"
		port_range = "80"
		target = "${google_compute_target_http_proxy.foobar1.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar1" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar2" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_backend_service" "foobar" {
		name = "%s"
		health_checks = ["${google_compute_http_health_check.zero.self_link}"]
	}

	resource "google_compute_http_health_check" "zero" {
		name = "%s"
		request_path = "/"
		check_interval_sec = 1
		timeout_sec = 1
	}

	resource "google_compute_url_map" "foobar" {
		name = "%s"
		default_service = "${google_compute_backend_service.foobar.self_link}"
		host_rule {
			hosts = ["mysite.com", "myothersite.com"]
			path_matcher = "boop"
		}
		path_matcher {
			default_service = "${google_compute_backend_service.foobar.self_link}"
			name = "boop"
			path_rule {
				paths = ["/*"]
				service = "${google_compute_backend_service.foobar.self_link}"
			}
		}
		test {
			host = "mysite.com"
			path = "/*"
			service = "${google_compute_backend_service.foobar.self_link}"
		}
	}`, fr, proxy1, proxy2, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_labels(fr, proxy1, proxy2, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
	resource "google_compute_global_forwarding_rule" "foobar" {
		description = "Resource created for Terraform acceptance testing"
		ip_protocol = "TCP"
		name = "%s"
		port_range = "80"
		target = "${google_compute_target_http_proxy.foobar1.self_link}"

	    labels {
	        my-label = "my-label-value"
	        my-second-label = "my-second-label-value"
	    }
	}

	resource "google_compute_target_http_proxy" "foobar1" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar2" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_backend_service" "foobar" {
		name = "%s"
		health_checks = ["${google_compute_http_health_check.zero.self_link}"]
	}

	resource "google_compute_http_health_check" "zero" {
		name = "%s"
		request_path = "/"
		check_interval_sec = 1
		timeout_sec = 1
	}

	resource "google_compute_url_map" "foobar" {
		name = "%s"
		default_service = "${google_compute_backend_service.foobar.self_link}"
		host_rule {
			hosts = ["mysite.com", "myothersite.com"]
			path_matcher = "boop"
		}
		path_matcher {
			default_service = "${google_compute_backend_service.foobar.self_link}"
			name = "boop"
			path_rule {
				paths = ["/*"]
				service = "${google_compute_backend_service.foobar.self_link}"
			}
		}
		test {
			host = "mysite.com"
			path = "/*"
			service = "${google_compute_backend_service.foobar.self_link}"
		}
	}`, fr, proxy1, proxy2, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_labelsUpdated(fr, proxy1, proxy2, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
	resource "google_compute_global_forwarding_rule" "foobar" {
		description = "Resource created for Terraform acceptance testing"
		ip_protocol = "TCP"
		name = "%s"
		port_range = "80"
		target = "${google_compute_target_http_proxy.foobar1.self_link}"

	    labels {
	        my-label = "my-label-value"
	        my-third-label = "my-third-label-value"
	    }
	}

	resource "google_compute_target_http_proxy" "foobar1" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar2" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_backend_service" "foobar" {
		name = "%s"
		health_checks = ["${google_compute_http_health_check.zero.self_link}"]
	}

	resource "google_compute_http_health_check" "zero" {
		name = "%s"
		request_path = "/"
		check_interval_sec = 1
		timeout_sec = 1
	}

	resource "google_compute_url_map" "foobar" {
		name = "%s"
		default_service = "${google_compute_backend_service.foobar.self_link}"
		host_rule {
			hosts = ["mysite.com", "myothersite.com"]
			path_matcher = "boop"
		}
		path_matcher {
			default_service = "${google_compute_backend_service.foobar.self_link}"
			name = "boop"
			path_rule {
				paths = ["/*"]
				service = "${google_compute_backend_service.foobar.self_link}"
			}
		}
		test {
			host = "mysite.com"
			path = "/*"
			service = "${google_compute_backend_service.foobar.self_link}"
		}
	}`, fr, proxy1, proxy2, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_basic2(fr, proxy1, proxy2, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
	resource "google_compute_global_forwarding_rule" "foobar" {
		description = "Resource created for Terraform acceptance testing"
		ip_protocol = "TCP"
		name = "%s"
		port_range = "80"
		target = "${google_compute_target_http_proxy.foobar2.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar1" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar2" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_backend_service" "foobar" {
		name = "%s"
		health_checks = ["${google_compute_http_health_check.zero.self_link}"]
	}

	resource "google_compute_http_health_check" "zero" {
		name = "%s"
		request_path = "/"
		check_interval_sec = 1
		timeout_sec = 1
	}

	resource "google_compute_url_map" "foobar" {
		name = "%s"
		default_service = "${google_compute_backend_service.foobar.self_link}"
		host_rule {
			hosts = ["mysite.com", "myothersite.com"]
			path_matcher = "boop"
		}
		path_matcher {
			default_service = "${google_compute_backend_service.foobar.self_link}"
			name = "boop"
			path_rule {
				paths = ["/*"]
				service = "${google_compute_backend_service.foobar.self_link}"
			}
		}
		test {
			host = "mysite.com"
			path = "/*"
			service = "${google_compute_backend_service.foobar.self_link}"
		}
	}`, fr, proxy1, proxy2, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_ipv6(fr, proxy1, proxy2, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
	resource "google_compute_global_forwarding_rule" "foobar" {
		description = "Resource created for Terraform acceptance testing"
		ip_protocol = "TCP"
		name = "%s"
		port_range = "80"
		target = "${google_compute_target_http_proxy.foobar1.self_link}"
		ip_version = "IPV6"
	}

	resource "google_compute_target_http_proxy" "foobar1" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_target_http_proxy" "foobar2" {
		description = "Resource created for Terraform acceptance testing"
		name = "%s"
		url_map = "${google_compute_url_map.foobar.self_link}"
	}

	resource "google_compute_backend_service" "foobar" {
		name = "%s"
		health_checks = ["${google_compute_http_health_check.zero.self_link}"]
	}

	resource "google_compute_http_health_check" "zero" {
		name = "%s"
		request_path = "/"
		check_interval_sec = 1
		timeout_sec = 1
	}

	resource "google_compute_url_map" "foobar" {
		name = "%s"
		default_service = "${google_compute_backend_service.foobar.self_link}"
		host_rule {
			hosts = ["mysite.com", "myothersite.com"]
			path_matcher = "boop"
		}
		path_matcher {
			default_service = "${google_compute_backend_service.foobar.self_link}"
			name = "boop"
			path_rule {
				paths = ["/*"]
				service = "${google_compute_backend_service.foobar.self_link}"
			}
		}
		test {
			host = "mysite.com"
			path = "/*"
			service = "${google_compute_backend_service.foobar.self_link}"
		}
	}`, fr, proxy1, proxy2, backend, hc, urlmap)
}
