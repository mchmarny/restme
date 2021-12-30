# GCS bucket to store exported logs
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

# Sink to drain Cloud Logs to the GCS bucket 
resource "google_logging_project_sink" "run_log_sink" {
  name                   = "${var.name}-revision-sink"
  description            = "Cloud Run logs"
  destination            = "storage.googleapis.com/${google_storage_bucket.logs_bucket.name}"
  filter                 = "resource.type = \"cloud_run_revision\" AND severity>=DEFAULT"
  unique_writer_identity = false
}

# IAM role binding to allow sink to write to GCS 
resource "google_project_iam_binding" "log_writer" {
  project = var.project_id
  role    = "roles/storage.objectCreator"

  members = [
    google_logging_project_sink.run_log_sink.writer_identity,
  ]

  lifecycle {
    ignore_changes = all
  }
}