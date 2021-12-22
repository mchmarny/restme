output "project_number" {
  value = data.google_project.project.number
}

output "external_url" {
  value = "${module.lb-http.external_ip} >> https://${var.name}.${var.domain}/"
}

output "workload_identity_pool_provider_id" {
  value = google_iam_workload_identity_pool_provider.github_provider.name
}

output "publisher_account_email" {
  value = google_service_account.publisher_service_account.email
}
