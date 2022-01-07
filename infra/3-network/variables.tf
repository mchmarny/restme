# List of variables which can be provided ar runtime to override the specified defaults 

variable "project_id" {
  description = "GCP Project ID"
  type        = string
  nullable    = false
}

variable "name" {
  description = "Service name"
  type        = string
  default     = "restme"
}

variable "domain" {
  description = "Domain name"
  type        = string
  nullable    = false
}

variable "regions" {
  description = "list of GPC regions to deploy to"
  type        = list(any)
  default     = ["us-west1", "europe-west1", "asia-east1"]
}

variable "alert_email" {
  type        = string
  description = "Email address to which alerts will be sent"
  nullable    = false
}
