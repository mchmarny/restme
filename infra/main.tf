# List of GCP APIs to enable in this project

resource "google_project_service" "compute" {
  project = var.project_id
  service = "compute.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "containerregistry" {
  project = var.project_id
  service = "containerregistry.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "dns" {
  project = var.project_id
  service = "dns.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "iam" {
  project = var.project_id
  service = "iam.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "iamcredentials" {
  project = var.project_id
  service = "iamcredentials.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "logging" {
  project = var.project_id
  service = "logging.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "monitoring" {
  project = var.project_id
  service = "monitoring.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "run" {
  project = var.project_id
  service = "run.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "secretmanager" {
  project = var.project_id
  service = "secretmanager.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "servicecontrol" {
  project = var.project_id
  service = "servicecontrol.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "servicemanagement" {
  project = var.project_id
  service = "servicemanagement.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "servicenetworking" {
  project = var.project_id
  service = "servicenetworking.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "stackdriver" {
  project = var.project_id
  service = "stackdriver.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "storage_api" {
  project = var.project_id
  service = "storage-api.googleapis.com"

  disable_on_destroy = false
}
resource "google_project_service" "storage_component" {
  project = var.project_id
  service = "storage-component.googleapis.com"

  disable_on_destroy = false
}