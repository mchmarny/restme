# List of outputs from each terraform apply 

output "PROJECT_ID" {
  value       = data.google_project.project.name
  description = "Project ID to use in Auth action for GCP in GitHub."
}

output "SERVICE_ACCOUNT" {
  value       = google_service_account.github_actions_user.email
  description = "Service account to use in GitHub Action for federated auth."
}

output "IDENTITY_PROVIDER" {
  value       = google_iam_workload_identity_pool_provider.github_provider.name
  description = "Provider ID to use in Auth action for GCP in GitHub."
}


