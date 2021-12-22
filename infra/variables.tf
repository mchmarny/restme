variable "org_id" {
  description = "GCP Org ID"
  type        = string
  default     = "110572507568"
}

variable "project_id" {
  description = "GCP Project ID"
  type        = string
  default     = "cloudy-labs"
}

variable "name" {
  description = "name prefix for resources"
  type        = string
  default     = "restme"
}

variable "domain" {
  description = "Domain name"
  type        = string
  default     = "cloudylab.dev"
}

variable "image" {
  description = "container image to deploy"
  type        = string
  default     = "gcr.io/cloudy-labs/restme:v0.6.27"
}

variable "regions" {
  description = "list of GPC regions to deploy to"
  type        = list(any)
  default     = ["us-west1", "europe-west1", "asia-east1"]
}

variable "registry_region" {
  type    = string
  default = "us-west1"
}

variable "limits" {
  type        = map(string)
  description = "Resource limits to the container"
  default = {
    cpu    = "1000m"
    memory = "512Mi"
  }
}

variable "ports" {
  type = object({
    name = string
    port = number
  })
  description = "Port which the container listens to (http1 or h2c)"
  default = {
    name = "http1"
    port = 8080
  }
}

variable "container_concurrency" {
  type        = number
  description = "Concurrent request limits to the service"
  default     = 80
}

variable "request_timeout" {
  type        = number
  description = "Timeout for each request in seconds"
  default     = 120
}

variable "max_scale" {
  type        = string
  description = "Maximum number of service instance annotation"
  default     = "10"
}

variable "api_key" {
  type        = string
  description = "API key version data"
  default     = "test-value"
}


variable "alert_email" {
  type        = string
  description = "Email address to which send alerts"
  default     = "mark+alert@chmarny.com"
}




