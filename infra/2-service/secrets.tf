# API Key Secret 
resource "google_secret_manager_secret" "secret_api_key" {
  secret_id = "${var.name}-api-key"

  labels = {
    label = "api-key"
  }

  replication {
    automatic = true
  }

  depends_on = [
    google_project_service.default["secretmanager.googleapis.com"],
  ]
}

data "template_file" "config" {
  template = file("../../configs/dev.json")
}

# API Key Secret version (holds data)
resource "google_secret_manager_secret_version" "secret_api_key_version" {
  secret = google_secret_manager_secret.secret_api_key.name

  secret_data = data.template_file.config.rendered
}

