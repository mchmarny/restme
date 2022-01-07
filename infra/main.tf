# List of GCP APIs to enable in this project

locals {
  services = [
    "compute.googleapis.com",
    "containerregistry.googleapis.com",
    "dns.googleapis.com",
    "iam.googleapis.com",
    "iamcredentials.googleapis.com",
    "logging.googleapis.com",
    "monitoring.googleapis.com",
    "run.googleapis.com",
    "secretmanager.googleapis.com",
    "servicecontrol.googleapis.com",
    "servicemanagement.googleapis.com",
    "servicenetworking.googleapis.com",
    "stackdriver.googleapis.com",
    "storage-api.googleapis.com",
  ]
}


resource "google_project_service" "default" {
  for_each = toset(local.services)

  project = var.project_id
  service = each.value

  timeouts {
    create = "10m"
  }
  
  disable_on_destroy = false
}