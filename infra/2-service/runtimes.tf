
# Service Account under which the Cloud Run services will run
resource "google_service_account" "runner_service_account" {
  account_id   = "${var.name}-runner"
  display_name = "${var.name}-runner"
}

# Policy to allow access to secrets 
data "google_iam_policy" "secret_reader" {
  binding {
    role = "roles/secretmanager.secretAccessor"

    members = [
      "serviceAccount:${google_service_account.runner_service_account.email}",
    ]
  }
}

# Binding of the secret access policy to the service account under which 
# Cloud Run services is running
resource "google_secret_manager_secret_iam_policy" "api_key_secret_access_policy" {
  project     = var.project_id
  secret_id   = google_secret_manager_secret.secret_api_key.secret_id
  policy_data = data.google_iam_policy.secret_reader.policy_data
}


# App Cloud Run service 
resource "google_cloud_run_service" "default" {
  for_each = toset(var.regions)

  name                       = "${var.name}--${each.value}"
  location                   = each.value
  project                    = var.project_id
  autogenerate_revision_name = true

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/${var.image}:${data.template_file.version.rendered}"
        volume_mounts {
          name       = "config-secret"
          mount_path = "/secrets"
        }
        ports {
          name           = "http1"
          container_port = 8080
        }
        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
        }
        env {
          name  = "ADDRESS"
          value = ":8080"
        }
        env {
          name  = "LOG_LEVEL"
          value = var.log_level
        }
        env {
          name  = "CONFIG"
          value = "/secrets/${var.name}"
        }
      }
      volumes {
        name = "config-secret"
        secret {
          secret_name = google_secret_manager_secret.secret_api_key.secret_id
          items {
            key  = var.secret_version
            path = var.name
          }
        }
      }

      container_concurrency = 80
      timeout_seconds       = 120
      service_account_name  = google_service_account.runner_service_account.email
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "10"
      }
    }
  }

  metadata {
    annotations = {
      "run.googleapis.com/client-name" = "terraform"
      "run.googleapis.com/ingress"     = "internal-and-cloud-load-balancing"
      # all, internal, internal-and-cloud-load-balancing
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  depends_on = [google_secret_manager_secret_version.secret_api_key_version]
}

resource "google_cloud_run_service_iam_member" "public-access" {
  for_each = toset(var.regions)

  location = google_cloud_run_service.default[each.key].location
  project  = google_cloud_run_service.default[each.key].project
  service  = google_cloud_run_service.default[each.key].name
  role     = "roles/run.invoker"
  member   = "allUsers"
}


