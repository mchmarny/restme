# restme

Collection of REST services

## services 

* [Request](#v1resource) - client request, headers and environment variables 
* [Resources](#v1resource) - host info, and RAM/CPU resources 

### /v1/request

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

### /v1/echo 

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

## deployment 

Initialize:

```sh
terraform init
```

Show plan

```sh
terraform plan -var=name=restme -var=project_id=cloudy-lab
```

Apply

```sh
terraform apply -var=name=restme -var=project_id=cloudy-lab
```

The output will be the LB IP 

## clean up

```sh
terraform destroy -var=name=restme -var=project_id=cloudy-lab
```

