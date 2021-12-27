# Cloud Armor policies 
# TODO: add cve-canary and throttle policies from infra/patch/policy 
#       when they become available in google-beta provider

resource "google_compute_security_policy" "policy" {
  name = "${var.name}-policy"

  lifecycle {
    ignore_changes = all
  }


  rule {
    action   = "deny(403)"
    priority = "1000"
    match {
      versioned_expr = "SRC_IPS_V1"
      config {
        src_ip_ranges = ["9.9.9.0/24"]
      }
    }
    description = "Deny access to IPs in 9.9.9.0/24"
  }

  rule {
    action   = "allow"
    priority = "2147483647"
    match {
      versioned_expr = "SRC_IPS_V1"
      config {
        src_ip_ranges = ["*"]
      }
    }
    description = "default rule"
  }
}