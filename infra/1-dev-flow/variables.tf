# List of variables which can be provided ar runtime to override the specified defaults 

variable "project_id" {
  description = "GCP Project ID"
  type        = string
  nullable    = false
}

variable "git_repo" {
  description = "GitHub Repo"
  type        = string
  nullable    = false
}


