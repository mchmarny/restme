provider "google" {
  project = var.project_id
}

provider "google-beta" {
  project = var.project_id
  region  = var.registry_region
}