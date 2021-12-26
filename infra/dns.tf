# DNS zone for specified domain to enable SSL certs
resource "google_dns_managed_zone" "dns_zone" {
  project     = var.project_id
  name        = replace(var.domain, ".dev", "")
  dns_name    = "${var.domain}."
  description = "${var.domain} DNS zone"

  dnssec_config {
    kind          = "dns#managedZoneDnsSecConfig"
    non_existence = "nsec3"
    state         = "on"

    default_key_specs {
      algorithm  = "ecdsap256sha256"
      key_length = 256
      key_type   = "keySigning"
      kind       = "dns#dnsKeySpec"
    }

    default_key_specs {
      algorithm  = "ecdsap256sha256"
      key_length = 256
      key_type   = "zoneSigning"
      kind       = "dns#dnsKeySpec"
    }
  }

  lifecycle {
    prevent_destroy = true
  }
}

# DNS A entry for the specified domain 
resource "google_dns_record_set" "a_app" {
  project      = var.project_id
  provider     = google-beta
  managed_zone = google_dns_managed_zone.dns_zone.name
  name         = "${var.name}.${var.domain}."
  type         = "A"
  rrdatas      = [google_compute_global_address.http_lb_address.address]
  ttl          = 60
}
