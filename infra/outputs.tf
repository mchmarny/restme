
output "external_url" {
  value = "${module.lb-http.external_ip} >> https://${var.domain}/"
}