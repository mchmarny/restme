# restme

Template to bootstrap a fully functional, multi-region, REST service on GCP with a developer release pipeline.

## Provisioned Infrastructure

* HTTPS Load Balancer configured with:
  * Custom domain
  * External IP
  * SSL certificate
  * Throttling and Canary CVE policies
  * Serverless NEGs
* Cloud Run service load balanced across 3 regions with:
  * Custom identity (service account)
  * Cloud Secrets variable
  * Custom capacity and autoscaling strategy 
  * Internal and Load Balancer traffic only trigger (no external access)
* Project configuration for:
  * Logging with GCS bucket sink 
  * Service uptime and SSL cert expiration alerts
  * Container registry 
  * Workload identity pool provider for GitHub Actions
  
## Development Workflow 

* Local test, lint, and validate actions using Makefile
* GitHub Actions for:
  * Test workflow on each PR
  * Image publishing to GCR and signing workflow on git tag


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

### Deployment

1. Update [terraform variables](infra/variables.tf) with your project ID, domain name, email address etc.

2. Authenticate to GCP using `gcloud`:

```shell
gcloud auth application-default login
```

3. Initialize terraform:

```shell
terraform init
```

4. Apply the configuration

> Review the displayed plan returned by terraform answering `yes`

```sh
terraform apply
```

> This will prompt you to provide values for each one of the defined variables. Alternatively, you can either define an environment file (e.g. `terraform.tfvars` in the same directory or pass the `-var-file` flag. (e.g. `terraform plan -var-file="envs/prod.tfvars"`).

After the resources are applied to your GCP project, terraform wil return: 

* `external_url` (IP and HTTPS address) 
* `publisher_account_email` (Service account `<name>-publisher@<project>.iam.gserviceaccount.com`)
* workload_identity_pool_provider_id (`projects/<project_number>/locations/global/workloadIdentityPools/github-id-pool-restme/providers/github-provider-restme`)

> I couldn't figure out how to apply the throttling and canary CVE policies in Cloud Armor using terraform. To apply these rules you will have to execute the [infra/patch/policy](infra/patch/policy) script.

5. Validate the deployment 

Use `curl` to access the value of `external_url` returned by the `terraform apply` command

### GitHub project configuration 

Create [GitHub secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) in your project for: 

* `PROJECT_ID` with the ID (not name) of your GCP project
* `SERVICE_ACCOUNT` with value of `publisher_account_email` returned by `terraform apply` command
* `WORKLOAD_IDENTITY_PROVIDER` with value of `workload_identity_pool_provider_id` also returned by `terraform apply` command

Now, whenever you run `make tag` locally, the `Publish` action will be triggered on GitHub and a new image will be published to GCR in your project. 

## Clean up

```sh
terraform destroy
```

