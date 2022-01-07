# List of GCP APIs to enable in this project
locals {
  services = [
    "servicecontrol.googleapis.com",
    "containerregistry.googleapis.com",
    "iam.googleapis.com",
    "iamcredentials.googleapis.com",
    "servicemanagement.googleapis.com",
    "storage-api.googleapis.com",
  ]
}

# Data source to access GCP project metadata 
data "google_project" "project" {}


resource "google_project_service" "default" {
  for_each = toset(local.services)

  project = var.project_id
  service = each.value

  disable_on_destroy = false
}