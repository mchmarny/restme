# restme

Collection of REST services

## services 

* [Request](#v1resource) - headers and environment variables 
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
      "GIN_MODE": "release",
      "HOME": "/home/nonroot",
      "HOSTNAME": "restme-86888bf66-9khgb",
      "KO_DATA_PATH": "/var/run/ko",
      "KUBERNETES_PORT": "tcp://10.96.0.1:443",
      "KUBERNETES_PORT_443_TCP": "tcp://10.96.0.1:443",
      "KUBERNETES_PORT_443_TCP_ADDR": "10.96.0.1",
      "KUBERNETES_PORT_443_TCP_PORT": "443",
      "KUBERNETES_PORT_443_TCP_PROTO": "tcp",
      "KUBERNETES_SERVICE_HOST": "10.96.0.1",
      "KUBERNETES_SERVICE_PORT": "443",
      "KUBERNETES_SERVICE_PORT_HTTPS": "443",
      "LOG_JSON": "true",
      "LOG_LEVEL": "debug",
      "PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/ko-app",
      "RESTME_PORT": "tcp://10.96.125.65:80",
      "RESTME_PORT_80_TCP": "tcp://10.96.125.65:80",
      "RESTME_PORT_80_TCP_ADDR": "10.96.125.65",
      "RESTME_PORT_80_TCP_PORT": "80",
      "RESTME_PORT_80_TCP_PROTO": "tcp",
      "RESTME_SERVICE_HOST": "10.96.125.65",
      "RESTME_SERVICE_PORT": "80",
      "SSL_CERT_FILE": "/etc/ssl/certs/ca-certificates.crt"
  }
}
```

### /v1/resource

```json
{
  "request": {
      "host": "172.18.0.3:30080",
      "id": "2996a496-b0dd-11eb-acc7-3e6bed2cd376",
      "method": "GET",
      "path": "/v1/resource",
      "protocol": "HTTP/1.1",
      "time": "2021-05-09T15:42:29.693151695Z",
      "version": "v0.3.1"
  },
  "host": {
      "hostname": "restme-86888bf66-9khgb",
      "uptime": 275,
      "boot_time": 1620574675,
      "processes": 1,
      "os": "linux",
      "platform": "debian",
      "platform_family": "debian",
      "platform_version": "10.9",
      "kernel_version": "5.4.0-1046-azure",
      "kernel_architecture": "x86_64",
      "host_id": "dbb592b7-88b8-104c-a512-b1ef05c7203a"
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
  "request": {
      "host": "172.18.0.3:30080",
      "id": "2a2e9f81-b0dd-11eb-acc7-3e6bed2cd376",
      "method": "GET",
      "path": "/v1/load/5s",
      "protocol": "HTTP/1.1",
      "time": "2021-05-09T15:42:30.689167384Z",
      "version": "v0.3.1"
  },
  "result": {
      "cores": 2,
      "start": 1620574950,
      "end": 1620574955,
      "operations": 1467302,
      "duration": "5s"
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