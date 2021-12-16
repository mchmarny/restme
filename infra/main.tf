provider "google" {
  project = var.project_id
}

resource "google_service_account" "service_account" {
  account_id   = "${var.name}-runner"
  display_name = "Service Account to run ${var.name}"
}

module "lb-http" {
  source            = "GoogleCloudPlatform/lb-http/google//modules/serverless_negs"
  version           = "~> 5.1"

  project = var.project_id
  name    = var.name

  ssl                             = true
  managed_ssl_certificate_domains = [var.domain]
  https_redirect                  = true

  backends = {
    default = {
      description             = null
      enable_cdn              = false
      security_policy         = null
      custom_request_headers  = null
      custom_response_headers = null

      groups = [
        for neg in google_compute_region_network_endpoint_group.serverless_neg:
        {
          group = neg.id
        }
      ]

      iap_config = {
        enable               = false
        oauth2_client_id     = ""
        oauth2_client_secret = ""
      }

      log_config = {
        enable      = false
        sample_rate = null
      }

    }
  }
}

resource "google_compute_region_network_endpoint_group" "serverless_neg" {
  for_each = toset(var.regions)

  name                  = "${var.name}--neg--${each.key}"
  network_endpoint_type = "SERVERLESS"
  region                = google_cloud_run_service.default[each.key].location
  cloud_run {
    service = google_cloud_run_service.default[each.key].name
  }
}

resource "google_cloud_run_service" "default" {
  for_each = toset(var.regions)

  name     = "${var.name}--${each.value}"
  location = each.value
  project  = var.project_id

  autogenerate_revision_name = true

  template {
    spec {
      containers {
        image = var.image
        resources {
          limits = {
            memory = var.memory
            cpu    = var.cpu
          }
        }
        env {
          name = "ADDRESS"
          value = ":8080"
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
      container_concurrency = 80
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


