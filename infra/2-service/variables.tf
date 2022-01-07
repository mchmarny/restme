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

variable "image" {
  description = "container image to deploy"
  type        = string
  nullable    = false
}

variable "regions" {
  description = "list of GPC regions to deploy to"
  type        = list(any)
  default     = ["us-west1", "europe-west1", "asia-east1"]
}

variable "log_level" {
  type        = string
  description = "Level of logging to use in the container (e.g. panic, fatal, error, warn, info, debug, trace)"
  default     = "info"
}

variable "service_limits" {
  type        = map(string)
  description = "Resource limits to the container"
  default = {
    cpu    = "1000m"
    memory = "512Mi"
  }
}

variable "service_ports" {
  type = object({
    name = string
    port = number
  })
  description = "Port which the container listens to (http1 or h2c)"
  default = {
    name = "http1"
    port = 8080
  }
  validation {
    condition     = var.service_ports.port > 1023
    error_message = "Ports below 1024 can be only opened by root."
  }
}

variable "service_concurrency" {
  type        = number
  description = "Concurrent request limits to the service"
  default     = 80
  validation {
    condition     = var.service_concurrency >= 1 && var.service_concurrency <= 1000
    error_message = "Number of requests in Cloud Run that can be processed simultaneously by a given container instance has to be between 1 and 1000."
  }
}

variable "service_timeout" {
  type        = number
  description = "Timeout for each request in seconds"
  default     = 120
  validation {
    condition     = var.service_timeout > 0
    error_message = "Request timeout has to be greater than 0."
  }
}

variable "service_max_scale" {
  type        = string
  description = "Maximum number of service instance annotation"
  default     = "10"
}

variable "api_key" {
  type        = string
  description = "Secret version data for API key"
  default     = "test-value"
}

