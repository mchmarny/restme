terraform {
  backend "gcs" {
    # GCS bucket where terrafrom state will be saved
    # Must exists before 1st terraform init
    bucket = "cloudylab-terraform-state"
    prefix = "prod"
  }
}