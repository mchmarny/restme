# List of outputs from each terraform apply 

output "external_ip" {
  value       = module.lb-http.external_ip
  description = "Resulting IP on LB that is routing traffic to Cloud Run services."
}

output "external_url" {
  value       = "https://${var.name}.${var.domain}/"
  description = "Resulting URL on LB that is routing traffic to Cloud Run services."
}