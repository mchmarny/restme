terraform {
  backend "gcs" {
    # GCS bucket where terrafrom state will be saved
    # Must exists before 1st terraform init
    # gsutil mb -p cloudy-lab2 -c STANDARD -l US-WEST1 -b on gs://restme-terraform-state
    bucket = "restme-terraform-state"
    prefix = "prod"
  }
}