variable "project_id" {
  description = "GCP Project ID"
  type        = string
  default     = "cloudy-lab"
}

variable "name" {
  description = "name prefix for resources"
  type        = string
  default     = "restme"
}

variable "image" {
  description = "container image to deploy"
  type        = string
  default     = "gcr.io/cloudy-lab/restme"
}

variable "regions" {
  description = "list of GPC regions to deploy to"
  type        = list
  default     = ["us-west1", "europe-west1", "asia-east1"]
}