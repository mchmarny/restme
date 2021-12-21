resource "google_storage_bucket" "logs_bucket" {
  project                     = var.project_id
  name                        = "${var.name}-logs"
  location                    = "US"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  lifecycle_rule {
    condition {
      age = "59"
    }
    action {
      type = "Delete"
    }
  }

  labels = {
    data_source = "cloud-run"
  }
}

resource "google_logging_project_sink" "run-log-sink" {
  name        = "${var.name}-sink"
  description = "Cloud Run logs"
  destination = "storage.googleapis.com/${google_storage_bucket.logs_bucket.name}"
  filter      = "resource.type = \"cloud_run_revision\" AND severity>=DEFAULT"

  unique_writer_identity = true
}

resource "google_project_iam_binding" "log-writer" {
  project = var.project_id
  role    = "roles/storage.objectCreator"

  members = [
    google_logging_project_sink.run-log-sink.writer_identity,
  ]
}