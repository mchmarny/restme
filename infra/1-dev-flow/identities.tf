locals {
  # List of roles that will be assigned to the pulbisher service account
  publisher_roles = toset([
    "roles/storage.objectCreator",
    "roles/storage.objectViewer",
  ])
}

# GCR registry (gcr.io)
resource "google_container_registry" "registry" {
  project = var.project_id

  depends_on = [
    google_project_service.default["containerregistry.googleapis.com"],
  ]
}

# Service account to be used for federated auth to publish to GCR
resource "google_service_account" "github_actions_user" {
  account_id   = "github-actions-user"
  display_name = "Service Account impersonated in GitHub Actions"
}

# Role binding to allow publisher to publish images
resource "google_project_iam_member" "github_actions_user_storage_role_binding" {
  for_each = local.publisher_roles
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.github_actions_user.email}"
}

# Default GCR bucket policy in GCS
data "google_iam_policy" "gcr_bucket_policy" {
  binding {
    role = "roles/storage.legacyBucketReader"
    members = [
      "serviceAccount:${google_service_account.github_actions_user.email}",
    ]
  }
}

# Assignment of the default GCR bucket policy in GCS to the registry
resource "google_storage_bucket_iam_policy" "default_gcr_policy" {
  bucket      = google_container_registry.registry.id
  policy_data = data.google_iam_policy.gcr_bucket_policy.policy_data
}


# Identiy pool for GitHub action based identity's access to Google Cloud resources
resource "google_iam_workload_identity_pool" "github_pool" {
  provider                  = google-beta
  workload_identity_pool_id = "github-pool"
}

# Configuration for GitHub identiy provider
resource "google_iam_workload_identity_pool_provider" "github_provider" {
  provider                           = google-beta
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider"
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.aud"        = "assertion.aud"
    "attribute.actor"      = "assertion.actor"
    "attribute.repository" = "assertion.repository"
  }
  oidc {
    issuer_uri        = "https://token.actions.githubusercontent.com"
    allowed_audiences = []
  }
}

# IAM policy bindings to the service account resources created by GitHub identify
resource "google_service_account_iam_member" "pool_impersonation" {
  provider           = google-beta
  service_account_id = google_service_account.github_actions_user.id
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_pool.name}/attribute.repository/${var.git_repo}"
}
