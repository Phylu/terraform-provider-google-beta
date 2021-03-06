// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeTargetTcpProxy_targetTcpProxyBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeTargetTcpProxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetTcpProxy_targetTcpProxyBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_target_tcp_proxy.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeTargetTcpProxy_targetTcpProxyBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_target_tcp_proxy" "default" {
  name            = "test-proxy-%s"
  backend_service = "${google_compute_backend_service.default.self_link}"
}

resource "google_compute_backend_service" "default" {
  name          = "backend-service-%s"
  protocol      = "TCP"
  timeout_sec   = 10

  health_checks = ["${google_compute_health_check.default.self_link}"]
}

resource "google_compute_health_check" "default" {
  name               = "health-check-%s"
  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "443"
  }
}
`, val, val, val,
	)
}

func testAccCheckComputeTargetTcpProxyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_target_tcp_proxy" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/targetTcpProxies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeTargetTcpProxy still exists at %s", url)
		}
	}

	return nil
}
