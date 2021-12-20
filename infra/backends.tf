terraform {
  backend "gcs" {
    bucket  = "cloudy-labs-terraform-state"
    prefix  = "prod"
  }
}