
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

resource "google_project_iam_member" "publisher_agent_binding" {
  project = var.project_id
  role    = "roles/cloudbuild.serviceAgent"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_project_iam_member" "publisher_storage_create_binding" {
  project = var.project_id
  role    = "roles/storage.objectCreator"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_project_iam_member" "publisher_storage_view_binding" {
  project = var.project_id
  role    = "roles/storage.admin"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_storage_bucket" "builder_logs_bucket" {
  name          = "${data.google_project.project.number}-cloudbuild-logs"
  location      = "US"
  force_destroy = true
}

data "google_iam_policy" "admin" {
  binding {
    role = "roles/storage.admin"
    members = [
      "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com",
      "serviceAccount:${google_service_account.publisher_service_account.email}",
    ]
  }
}

resource "google_storage_bucket_iam_policy" "build_logs_bucket_policy" {
  bucket = google_storage_bucket.builder_logs_bucket.name
  policy_data = data.google_iam_policy.admin.policy_data
}


