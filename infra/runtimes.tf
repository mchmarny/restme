
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
          limits = var.limits
        }
        env {
          name  = "ADDRESS"
          value = ":${var.ports["port"]}"
        }
        env {
          name  = "IMAGE"
          value = var.image
        }
        env {
          name  = "GIN_MODE"
          value = "release"
        }
        env {
          name  = "LOG_LEVEL"
          value = "debug"
        }
        env {
          name = "API_KEY"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.secret_api_key.secret_id
              key  = "latest"
            }
          }
        }
      }

      container_concurrency = var.container_concurrency
      timeout_seconds       = var.request_timeout
      service_account_name  = google_service_account.runner_service_account.email
    }
  }

  metadata {
    labels = {
      terraformed = "true"
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [
    google_secret_manager_secret_iam_policy.api_key_secret_access_policy,
  ]
}

resource "google_cloud_run_service_iam_member" "public-access" {
  for_each = toset(var.regions)

  location = google_cloud_run_service.default[each.key].location
  project  = google_cloud_run_service.default[each.key].project
  service  = google_cloud_run_service.default[each.key].name
  role     = "roles/run.invoker"
  member   = "allUsers"
}


