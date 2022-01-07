# Cloud Run service to be deployed in each region
data "google_cloud_run_service" "default" {
    for_each = toset(var.regions)

    name = "${var.name}--${each.value}"
    location = each.value
}