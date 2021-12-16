variable "project_id" {
  description = "GCP Project ID"
  type        = string
  default     = "cloudy-lab"
}

variable "release" {
  description = "Git commit sha (e.g. git rev-parse --short HEAD)"
  type        = string
  default     = "1acc66a"
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

variable "domain" {
  description = "Domain for SSL cert"
  type        = string
  default     = "restme.cloudylab.dev"
}

variable "memory" {
  description = "Memory limit for container"
  type        = string
  default     = "512Mi"
}

variable "cpu" {
  description = "CPU limit for container"
  type        = string
  default     = "1000m"
}