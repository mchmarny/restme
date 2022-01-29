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

Go code exposing following sample services:

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
      "version": "v0.8.11"
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

1. In terminal, from the root of the project, first initialize terraform

> Note, this flow uses local terraform state, make sure you do not check that into source control and consider using persistent state provider like GCS

```shell
terraform -chdir=infra/1-dev-flow init
```

2. Apply the configuration

> When promoted for the `GCP Project ID`, enter your existing project ID (not the name), `GitHub Repo` name (in username/repo-name format), and confirm with `yes` the terraform displayed plan. Alternatively you can use either the command-line variables or a terraform variable file. More on that [here](https://www.terraform.io/language/values/variables).

```sh
terraform -chdir=infra/1-dev-flow apply
```

3. Create GitHub Secrets 

The result of the `apply` command above will look something like this: 

```shell
PROJECT_ID = "<id-of-your-project>"
SERVICE_ACCOUNT = "github-action-publisher@<id-of-your-project>.iam.gserviceaccount.com"
IDENTITY_PROVIDER = "projects/<project-number>/locations/global/workloadIdentityPools/github-pool/providers/github-provider"
```

Navigate to your GitHub project [secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) and create the following entries with the values returned by the apply command: 

* `PROJECT_ID`
* `SERVICE_ACCOUNT`
* `IDENTITY_PROVIDER`

4. Test GitHub Workflow 

Now, whenever you create a version tag (`v*`) in your repo, GitHub will run the [image-on-tag.yaml](.github/workflows/image-on-tag.yaml) action which will build, publish, and sign (using cosign) your container image in GCR. 

> This assumes you have not already created a git tag with this v0.0.1 version

```shell
git tag v0.0.1
git push origin v0.0.1
```

Once the action completes, you should see the new image in GCR (`gcr.io/<project_id>/restme:v0.0.1`).

## Cloud Run Service Deployment

> This assumes an already published image in GCR

1. Initialize terraform

> Note, this flow uses local terraform state, make sure you do not check that into source control and consider using persistent state provider like GCS

The service deployment step uses few new Terraform modules, so start by initializing the deployment. 

```shell
terraform -chdir=infra/2-service init --upgrade
```

2. Apply the configuration

This deployment will prompt for a lot of variables, you can create `variables.tf` with the following entries in `infra/2-service` folder to avoid these prompts. Edit these as necessary. 

```txt
project_id     = "your-project-id"
name           = "restme"
domain         = "your.domain.dev"
regions        = ["us-west1", "europe-west1", "asia-east1"]
image          = "restme"
secret_version = "latest"
log_level      = "info"
```

> Note, the domain must be something you can control DNS for as you will have to create an `A` entry to point to the `IP` in Terraform output for this step. 


```sh
terraform -chdir=infra/2-service apply -var-file=terraform.tfvars
```

The result of `apply` should be a list of Cloud Run services (URLs) for each one of the regions you deployed to, as well as the external IP and URL on the load balancer by which you can access these services. 

> Note, you will not be able to access the Cloud Run services directly as their ingress (trigger) is internal and Cloud load balancer only. 

```shell
cloud_run_services = toset([
  "https://restme--asia-east1-fr3j36toba-de.a.run.app",
  "https://restme--europe-west1-fr3j36toba-ew.a.run.app",
  "https://restme--us-west1-fr3j36toba-uw.a.run.app",
])
external_ip = "x.x.x.x"
external_url = "https://your.domain.dev/"
```

It will take a few min for the SSL certificate to be provisioned. As soon as the `apply` step completes, use the IP in `external_ip` to create an `A` record in your DNS to point to the domain in `external_url`. For example: 

Host Name: `demo` # the `your` portion of `your.domain.dev`
Type: `A`
TTL: `60`
Data: `x.x.x.x` # the actual IP returned by the above step

> Note, the SSL cert provisioning in the above action will take a few min. 

Assuming everything went OK, you should now be able to test the deployment by using `curl` to invoke the address returned by the `external_url` output

```shell
url https://your.domain.dev/
```

## Clean up

To clean up each of these deployments run the following command:

> Note, the project itself and the external IP address will not be deleted and all APIs enabled as part of these deployments will stay enabled. 

```sh
terraform -chdir=infra/2-service destroy
terraform -chdir=infra/1-dev-flow destroy
```

