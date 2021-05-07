# restme

Collection of REST services to user in platform validation.

## services 

* [Request](#v1resource)
* [Resources](#v1resources)
* [Load](#v1load)

### /v1/request

> Note, values obfuscated and generalized for illustration purposes

```json
{
  "request": {
    "id": "0cc38a00-ae86-11eb-a0c9-e6b1fa9e5990",
    "time": "2021-05-06T16:13:52.792423514Z",
    "uri": "/v1/request",
    "host": "app.domain.com",
    "method": "GET"
  },
  "headers": {
    "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "accept-encoding": "gzip, deflate, br",
    "accept-language": "en-us",
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15",
    "x-b3-spanid": "81f0***",
    "x-b3-traceid": "0ed5***",
    "x-forwarded-by": "proxy.domain.com",
    "x-forwarded-for": "*.*.*.*",
    "x-forwarded-info": "Nginx 1.*.0",
    "x-forwarded-proto": "https",
    "x-real-ip": "*.*.*.*"
  },
  "env_vars": {
    "APP0_PORT": "tcp://10.1.1.14:80",
    "APP0_PORT_80_TCP": "tcp://10.1.1.14:80",
    "APP0_PORT_80_TCP_ADDR": "10.1.1.14.132",
    "APP0_PORT_80_TCP_PORT": "80",
    "APP0_PORT_80_TCP_PROTO": "tcp",
    "APP0_SERVICE_HOST": "10.1.1.14",
    "APP0_SERVICE_PORT": "80",
    "APP0_SERVICE_PORT_WEB": "80",
    "HOME": "/home/nonroot",
    "HOSTNAME": "app-12345678-x9e8",
    "KUBERNETES_PORT": "tcp://192.168.0.1:443",
    "KUBERNETES_PORT_443_TCP": "tcp://192.168.0.1:443",
    "KUBERNETES_PORT_443_TCP_ADDR": "192.168.0.1",
    "KUBERNETES_PORT_443_TCP_PORT": "443",
    "KUBERNETES_PORT_443_TCP_PROTO": "tcp",
    "KUBERNETES_SERVICE_HOST": "192.168.0.1",
    "KUBERNETES_SERVICE_PORT": "443",
    "KUBERNETES_SERVICE_PORT_HTTPS": "443",
    "PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
  }
}
```

### /v1/resource

> Note, values obfuscated and generalized for illustration purposes

```json
{
  "request": {
    "id": "0c923055-ae86-11eb-a0c9-e6b1fa9e5990",
    "time": "2021-05-06T16:13:52.469001657Z",
    "uri": "/v1/resource",
    "host": "app.domain.com",
    "method": "GET"
  },
  "host": {
    "hostname": "app-12345678-w5h8b",
    "uptime": 237181,
    "boot_time": 1620080451,
    "processes": 1,
    "os": "linux",
    "platform": "debian",
    "platform_family": "debian",
    "platform_version": "10.9",
    "kernel_version": "5.4.77-7.el7pie",
    "kernel_architecture": "x86_64",
    "virtualization_system": "kvm",
    "virtualization_role": "host",
    "host_id": "046a9c8a-ae88-11eb-8529-0242ac130003"
  },
  "resources": {
    "ram": {
      "value": "376.3G",
      "context": "Source: OS process status, Size: 376.3G"
    },
    "cpu": {
      "value": 32,
      "context": "Source: OS process status"
    }
  },
  "limits": {
    "ram": {
      "value": "512M",
      "context": "source: /sys/fs/cgroup/memory/memory.limit_in_bytes, writable: false, size: 512M"
    },
    "cpu": {
      "value": 1,
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
  "id": "3684ab34-aed8-11eb-ba54-1e00d11edc71",
  "time": "2021-05-06T19:02:01.57696-07:00",
  "uri": "/v1/load?duration=5s",
  "host": "app.domain.com",
  "method": "GET"
 },
 "result": {
  "cores": 8,
  "start": 1620352921,
  "end": 1620352926,
  "operations": 8454324,
  "duration": "5s"
 }
}
```