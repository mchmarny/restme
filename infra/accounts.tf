
resource "google_service_account" "service_account" {
  account_id   = "${var.name}-runner"
  display_name = "${var.name}-runner"
}

