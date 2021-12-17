terraform {
  backend "gcs" {
    bucket = "restme-prod"
    prefix = "terraform/state"
  }
}