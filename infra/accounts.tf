
resource "google_service_account" "runner_service_account" {
  account_id   = "${var.name}-runner"
  display_name = "${var.name}-runner"
}

resource "google_service_account" "publisher_service_account" {
  account_id   = "${var.name}-publisher"
  display_name = "${var.name}-publisher"
}

resource "google_project_iam_member" "publisher_builder_binding" {
  project = var.project_id
  role    = "roles/cloudbuild.builds.editor"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_project_iam_member" "publisher_storage_binding" {
  project = var.project_id
  role    = "roles/storage.objectCreator"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}