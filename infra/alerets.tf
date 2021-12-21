resource "google_monitoring_notification_channel" "notification_channel" {
  project      = var.project_id
  display_name = "Notification Channel"
  type         = "email"
  labels = {
    email_address = var.alert_email
  }
}

resource "google_monitoring_uptime_check_config" "uptime_check" {
  display_name = "restme-uptime-check"
  timeout      = "60s"

  http_check {
    path         = "/"
    port         = "443"
    use_ssl      = true
    validate_ssl = true
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = var.project_id
      host       = var.domain
    }
  }
}

resource "google_monitoring_alert_policy" "ssl_cert_expire_alert_policy" {
  display_name = "SSL certificate expiring soon"
  combiner     = "OR"
  enabled      = true
  conditions {
    display_name = "SSL certificate expiring soon"
    condition_threshold {
      filter          = "metric.type=\"monitoring.googleapis.com/uptime_check/time_until_ssl_cert_expires\" AND resource.type=\"uptime_url\""
      duration        = "600s"
      comparison      = "COMPARISON_LT"
      threshold_value = 15
      trigger {
        count = 1
      }
      aggregations {
        alignment_period     = "1200s"
        per_series_aligner   = "ALIGN_NEXT_OLDER"
        cross_series_reducer = "REDUCE_MEAN"
        group_by_fields      = ["resource.label.*"]
      }
    }
  }

  user_labels = {
    uptime  = "ssl_cert_expiration"
    version = 1
  }

  notification_channels = [
    google_monitoring_notification_channel.notification_channel.id
  ]
}


resource "google_monitoring_alert_policy" "uptime_alert_policy" {
  display_name = "Uptime alert policy"
  combiner     = "OR"
  enabled      = true
  conditions {
    display_name = "Failure of uptime check"
    condition_threshold {
      filter          = "resource.type = \"uptime_url\" AND metric.type = \"monitoring.googleapis.com/uptime_check/check_passed\" AND metric.labels.check_id = \"${google_monitoring_uptime_check_config.uptime_check.id}\""
      duration        = "60s"
      comparison      = "COMPARISON_GT"
      threshold_value = 1
      trigger {
        count = 1
      }
      aggregations {
        alignment_period     = "1200s"
        per_series_aligner   = "ALIGN_NEXT_OLDER"
        cross_series_reducer = "REDUCE_COUNT_FALSE"
        group_by_fields = [
          "resource.label.project_id",
          "resource.label.host"
        ]
      }
    }
  }

  user_labels = {
    uptime  = "uptime_check"
    version = 1
  }

  notification_channels = [
    google_monitoring_notification_channel.notification_channel.id
  ]
}
