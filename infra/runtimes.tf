
resource "google_cloud_run_service" "default" {
  for_each = toset(var.regions)

  name                       = "${var.name}--${each.value}"
  location                   = each.value
  project                    = var.project_id
  autogenerate_revision_name = true

  template {
    spec {
      containers {
        image = var.image
        ports {
          name           = var.ports["name"]
          container_port = var.ports["port"]
        }
        resources {
          limits   = var.limits
        }
        env {
          name = "ADDRESS"
          value = ":${var.ports["port"]}"
        }
        env {
          name = "GIN_MODE"
          value = "release"
        }
        env {
          name = "LOG_LEVEL"
          value = "debug"
        }
      }
      container_concurrency = var.container_concurrency
      timeout_seconds       = var.request_timeout
      service_account_name  = google_service_account.service_account.email
    }
  }

  metadata {
    labels = {
      terraformed = "true"
      release     = "${var.release}"
    }
  }

  depends_on = [google_service_account.service_account]
}

resource "google_cloud_run_service_iam_member" "public-access" {
  for_each = toset(var.regions)

  location = google_cloud_run_service.default[each.key].location
  project  = google_cloud_run_service.default[each.key].project
  service  = google_cloud_run_service.default[each.key].name
  role     = "roles/run.invoker"
  member   = "allUsers"
}


