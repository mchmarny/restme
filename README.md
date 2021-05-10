# restme

Collection of REST services

## services 

* [Request](#v1resource) - client request, headers and environment variables 
* [Resources](#v1resource) - host info, and RAM/CPU resources 
* [Load](#v1load)

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

### /v1/resource

```json
{
  "host": {
      "hostname": "restme-86888bf66-9khgb",
      "os": "linux",
      ...
  },
  "resources": {
      "ram": {
          "value": "6.8G",
          "context": "Source: OS process status, Size: 6.8G"
      },
      "cpu": {
          "value": 2,
          "context": "Source: OS process status"
      }
  },
  "limits": {
      "ram": {
          "value": "8388608T",
          "context": "source: /sys/fs/cgroup/memory/memory.limit_in_bytes, writable: false, size: 8388608T"
      },
      "cpu": {
          "context": "source: /sys/fs/cgroup/cpu/cpu.cfs_quota_us, writable: false"
      }
  }
}
```

### /v1/load

> Note, uses all available cores at 100%. Use load duration parameter: `/v1/load?duration=5s` to specify the duration of the load generation across all cores. 

```json
{
  "result": {
      "cores": 2,
      "operations": 1467302,
      "duration": "5s"
      ...
  }
}
```

### /v1/echo 

Request:

```shell
curl -i \
  -H "Content-Type: application/json" \
  http://localhost:8080/v1/echo \
  -d '{ "on": $(shell date +%s), "msg": "hello?" }'
```

Response: 

```json
{ "on": 1620438323, "msg": "hello?" }
```