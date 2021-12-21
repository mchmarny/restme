resource "google_service_account" "publisher_service_account" {
  account_id   = "${var.name}-publisher"
  display_name = "${var.name}-publisher"
}

resource "google_project_iam_member" "publisher_builder_binding" {
  project = var.project_id
  role    = "roles/cloudbuild.builds.editor"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_project_iam_member" "publisher_viewer_binding" {
  project = var.project_id
  role    = "roles/viewer"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_project_iam_member" "publisher_storage_binding" {
  project = var.project_id
  role    = "roles/storage.objectCreator"
  member  = "serviceAccount:${google_service_account.publisher_service_account.email}"
}

resource "google_iam_workload_identity_pool" "github_pool" {
  provider                  = google-beta
  workload_identity_pool_id = "github-id-pool-${var.name}"
}

resource "google_iam_workload_identity_pool_provider" "github_provider" {
  provider                           = google-beta
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider-${var.name}"
  attribute_mapping = {
    "google.subject"  = "assertion.sub"
    "attribute.aud"   = "assertion.aud"
    "attribute.actor" = "assertion.actor"
  }
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

resource "google_service_account_iam_member" "pool_impersonation" {
  provider           = google-beta
  service_account_id = google_service_account.publisher_service_account.id
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_pool.name}/*"
}
