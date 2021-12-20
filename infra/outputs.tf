
output "external_url" {
  value = "${module.lb-http.external_ip} >> https://${var.domain}/"
}

output "workload_identity_pool_provider_id" {
  value = google_iam_workload_identity_pool_provider.github_provider.name
}