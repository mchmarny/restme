
resource "google_service_account" "runner_service_account" {
  account_id   = "${var.name}-runner"
  display_name = "${var.name}-runner"
}

resource "google_service_account" "publisher_service_account" {
  account_id   = "${var.name}-publisher"
  display_name = "${var.name}-publisher"
}

