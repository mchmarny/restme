# List of cloud run services created in terraform apply 

output "cloud_run_services" {
  value = toset([
    for s in google_cloud_run_service.default : s.status[0].url
  ])
}