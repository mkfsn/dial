# dial

A Go based network health check tool.

# Example

```bash
$ dial -host github.com -port 443 -timeout 5s -retry 3 -verbose
dialing to 'github.com:443'
dial to 'github.com:443' successfully
```

# Usage

```bash
Usage of dial:
  -delay string
    	duration to delay retry, default is 1s (default "1s")
  -host string
    	host to dial, e.g. 127.0.0.1
  -network string
    	network to dial, default is tcp. (default "tcp")
  -port int
    	port to dial, e.g. 80
  -retry int
    	times to retry, default is 0 - no retry
  -timeout string
    	dial timeout, default is 10s (default "10s")
  -verbose
    	default is off
```
