# Cloud Armor policies 
# TODO: add cve-canary and throttle policies from infra/patch/policy 
#       when they become available in google-beta provider

resource "google_compute_security_policy" "policy" {
  name = "${var.name}-policy"

  lifecycle {
    ignore_changes = all
  }


  rule {
    action      = "deny(403)"
    description = "CVE-2021-44228 and CVE-2021-45046"
    priority    = 950
    match {
      expr {
        expression = "evaluatePreconfiguredExpr('cve-canary')"
      }
    }
  }

  rule {
    action   = "deny(403)"
    priority = 1000
    match {
      versioned_expr = "SRC_IPS_V1"
      config {
        src_ip_ranges = ["9.9.9.0/24"]
      }
    }
    description = "Deny access to IPs in 9.9.9.0/24"
  }
  
  rule {
    action      = "deny(403)"
    description = "Common DNS sniffing targets"
    priority    = 1200
    match {
      expr {
        expression = "request.path.matches('/Autodiscover|/bin/|/ecp/|/owa/|/vendor/|/ReportServer|/_ignition')"
      }
    }
  }

  rule {
    action   = "allow"
    priority = 2147483647
    match {
      versioned_expr = "SRC_IPS_V1"
      config {
        src_ip_ranges = ["*"]
      }
    }
    description = "default rule"
  }

}
