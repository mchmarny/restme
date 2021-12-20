variable "project_id" {
  description = "GCP Project ID"
  type        = string
  default     = "cloudy-lab"
}

variable "release" {
  description = "Git commit sha (e.g. git rev-parse --short HEAD)"
  type        = string
  default     = "d0b54b3"
}

variable "name" {
  description = "name prefix for resources"
  type        = string
  default     = "restme"
}

variable "image" {
  description = "container image to deploy"
  type        = string
  default     = "gcr.io/cloudy-lab/restme:v0.6.18"
}

variable "regions" {
  description = "list of GPC regions to deploy to"
  type        = list
  default     = ["us-west1", "europe-west1", "asia-east1"]
}

variable "registry_region" {
  type    = string
  default = "us-west1"
}

variable "domain" {
  description = "Domain for SSL cert"
  type        = string
  default     = "restme.cloudylab.dev"
}

variable "limits" {
  type        = map(string)
  description = "Resource limits to the container"
  default     = {
    cpu = "1000m"
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

variable "api_key" {
  type        = string
  description = "API key version data"
  default     = "test-value"
}




