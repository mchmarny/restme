resource "google_secret_manager_secret" "secret_api_key" {
  secret_id = "${var.name}-api-key"

  labels = {
    label = "api-key"
  }

  replication {
    automatic = true
  }
}


resource "google_secret_manager_secret_version" "secret-api-key-version" {
  secret = google_secret_manager_secret.secret_api_key.id

  secret_data = var.api_key
}

data "google_iam_policy" "secret_reader" {
  binding {
    role = "roles/secretmanager.secretAccessor"

    members = [
      "serviceAccount:${google_service_account.runner_service_account.email}",
    ]
  }
}

resource "google_secret_manager_secret_iam_policy" "api_key_secret_access_policy" {
  project = var.project_id
  secret_id = google_secret_manager_secret.secret_api_key.secret_id
  policy_data = data.google_iam_policy.secret_reader.policy_data
}
