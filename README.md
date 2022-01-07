# restme

Template to bootstrap a fully functional, multi-region, REST service on GCP with a developer release pipeline.

## Provisioned

List of resources provisioned using this project 

### Development Flow 

* Container registry (GCR) to store images
* Service account to use for image publishing from GitHub to GCR
* Workload identity pool provider for GitHub Action workflows to access GCR without the needing to store the GCP credentials as long-lived GitHub secrets

### Cloud Run Service 

Cloud Run service provisioned into n-number of regions with: 
  
* Custom identity (service account)
* Custom capacity and autoscaling strategy 
* Accessible only via Internal and Load Balancer traffic (i.e. no external access)
* Secret Manager service with sample secret 
* Revision config using secret variable

### Cloud Load Balancer

HTTPS load balancer configured with:

  * External IP
  * SSL certificate
  * Cloud DNS service with A record for custom domain
  * Cloud Armor service with throttling and Canary CVE policies
  * Serverless NEGs for load balancer to point to Cloud Run service in each region
  * Service uptime and SSL cert expiration alerts
  
## REST Services

Go code exposing following routes:

* [Request info](#getv1requestinfo) - client request, headers and environment variables 
* [Echo message](#postv1echomessage) - simple echo message 

### GET /v1/request/info

> Note, values obfuscated and generalized for illustration purposes

```json
{
  "request": {
      "host": "172.18.0.3:30080",
      "id": "2994eb6e-b0dd-11eb-acc7-3e6bed2cd376",
      "method": "GET",
      "path": "/v1/request",
      "protocol": "HTTP/1.1",
      "time": "2021-05-09T15:42:29.682509631Z",
      "version": "v0.3.1"
  },
  "headers": {
      "accept": "*/*",
      "content-type": "application/json",
      "user-agent": "curl/7.68.0"
  },
  "env_vars": {
      "HOME": "/home/nonroot",
      "HOSTNAME": "restme-86888bf66-9khgb",
      ...
  }
}
```

### POST /v1/echo/message

Request:

```shell
curl -i \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ..." \
  http://localhost:8080/v1/echo \
  -d '{ "on": $(shell date +%s), "msg": "hello?" }'
```

Response: 

```json
{ "on": 1620438323, "msg": "hello?" }
```

## Deployment

### Prerequisites 

> Good how-to on using terraform with GCP is located [here](https://cloud.google.com/community/tutorials/getting-started-on-gcp-with-terraform)

* [terraform CLI](https://www.terraform.io/downloads)
* [GCP Project](https://cloud.google.com/resource-manager/docs/creating-managing-projects)
* [gcloud CLI](https://cloud.google.com/sdk/gcloud)
  * Make sure to authenticate to GCP using `gcloud auth application-default login`

## Development Flow Deployment

1. In terminal, from the root of the project, `cd infra/1-dev-flow`

2. Initialize terraform

> Note, this flow uses local terraform state, make sure you do not check that into source control and consider using persistent state provider like GCS

```shell
terraform init
```

3. Apply the configuration

> When promoted for the `GCP Project ID`, enter your existing project code and confirm with `yes` the terraform displayed plan

```sh
terraform apply
```

4. Create GitHub Secrets 

The result of the `apply` command above will look something like this: 

```shell
PROJECT_ID = "<id-of-your-project>"
SERVICE_ACCOUNT = "github-action-publisher@<id-of-your-project>.iam.gserviceaccount.com"
WORKLOAD_IDENTITY_PROVIDER = "projects/<number-of-your-project>/locations/global/workloadIdentityPools/github-pool/providers/github-provider"
```

Navigate to your GitHub project [secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) and create the following entries with the values returned by the apply command: 

* `PROJECT_ID`
* `SERVICE_ACCOUNT`
* `WORKLOAD_IDENTITY_PROVIDER`

Now, whenever you create a version tag (`v*`) in your repo, GitHub will run the [image-on-tag.yaml](.github/workflows/image-on-tag.yaml) action which will build, publish, and sign (using cosign) your container image in GCR. 

## Cloud Run Service Deployment

> This assumes an already published image in GCR

1. In terminal, from the root of the project, `cd infra/2-service`

2. Initialize terraform

> Note, this flow uses local terraform state, make sure you do not check that into source control and consider using persistent state provider like GCS

```shell
terraform init
```

3. Apply the configuration

> For this demo, the [variables.tf](infra/2-service/variables.tf) file includes a lot of default values (like for example the list of regions to deploy to). Edit these as necessary. 

```sh
terraform apply
```

> This will prompt you to provide values for `project_id` and `image` (this is the previously published container image in GCR, the value should be in `gcr.io/<your-project-id>/restme:<tag-version>` format). Alternatively, you can define an environment file (e.g. `terraform.tfvars` in the same directory or pass the `-var-file` flag. (e.g. `terraform plan -var-file="envs/prod.tfvars"`). When promoted to confirm the plan, type `yes`.

The result of `apply` should be a list of Cloud Run services (URLs) for each one of the regions you deployed to. 

```shell
cloud_run_services = toset([
  "https://restme--asia-east1-4qt7uwb6vq-de.a.run.app",
  "https://restme--europe-west1-4qt7uwb6vq-ew.a.run.app",
  "https://restme--us-west1-4qt7uwb6vq-uw.a.run.app",
])
```

> Note, you will not be able to access these services yet since we annotated each one of the Cloud Run services with ingress (trigger) by internal and Cloud load balancer only. If you won't be using Cloud Load balancer (next step), you can remove the `"run.googleapis.com/ingress"` annotation in [runtimes.tf](infra/2-service/runtimes.tf) file. 

```json
{
  ...
  metadata {
    labels = {
      terraformed = "true"
    }
    annotations = {
      "autoscaling.knative.dev/maxScale" = var.service_max_scale
      "run.googleapis.com/client-name"   = "terraform"
      "run.googleapis.com/ingress"       = "internal-and-cloud-load-balancing"
    }
  }
  ...
}
```

### Cloud Load Balancer Deployment

> This assumes an already deployed Cloud Run services 

1. In terminal, from the root of the project, `cd infra/3-network`

2. Initialize terraform

> Note, this flow uses local terraform state, make sure you do not check that into source control and consider using persistent state provider like GCS

```shell
terraform init
```

3. Apply the configuration

> For this demo, the [variables.tf](infra/3-network/variables.tf) file includes a lot of default values (like for example the list of regions to deploy to). Edit these as necessary.

```sh
terraform apply
```

> This will prompt you to provide values for `project_id` and `domain` (this is full domain that will be used in the SSL cert), and `alert_email` where the alerts will be sent. 

The result of `apply` should be the external IP address (`external_ip`) and the load balanced URL (`external_url`). Make sure the domain you provided in `apply` points to that IP. 

4. Apply manual policies 

The Google Cloud provider is still missing some of the new Cloud Armor policy types so to apply the throttling policy you will have to create it manually. 

```shell
patch/policy
```

> Note, the SSL cert provisioning in the above action will take a few min. 

Assuming everything went OK, you should now be able to test the deployment by using `curl` to invoke the address returned by the `external_url` output

```shell
url https://restme.<your-domain>.dev/
```

Should rerun something like this:

```json
{
    "routes": [
        "POST    /v1/echo/message",
        "GET     /v1/request/info"
    ]
}
```

## Clean up

To clean up each of these deployments run the following command:

> Note, the project itself and the external IP address will not be deleted and all APIs enabled as part of these deployments will stay enabled. 

```sh
terraform destroy
```

