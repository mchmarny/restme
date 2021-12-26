# List of outputs from each terraform apply 

output "external_ip" {
  value       = module.lb-http.external_ip
  description = "Resulting IP on LB that is routing traffic to Cloud Run services."
}

output "external_url" {
  value       = "https://${var.name}.${var.domain}/"
  description = "Resulting URL on LB that is routing traffic to Cloud Run services."
}

output "workload_identity_pool_provider_id" {
  value       = google_iam_workload_identity_pool_provider.github_provider.name
  description = "ID of the Identity provider to use in Auth action for GCP in GitHub."
}

output "publisher_account_email" {
  value       = google_service_account.publisher_service_account.email
  description = "Service account to use in GitHub Action for federated auth."
}
