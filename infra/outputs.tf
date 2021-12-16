
output "url" {
  value = "${module.lb-http.external_ip} >> https://${var.domain}/"
}