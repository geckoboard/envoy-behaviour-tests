Small example app/config to understand how envoy behaves if an upstream
server takes longer than the `per_try_timeout` option to response, but
shorter than the overall timeout.

Try adding/removing the `hedge_policy` section in the route and watch how
you never get a response from the backend.

When you enable the hedging policy backend the client sees the response
from the first request:

```
❯ curl 127.0.0.1:9091/foo --verbose
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 9091 (#0)
> GET /foo HTTP/1.1
> Host: 127.0.0.1:9091
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< date: Thu, 12 Mar 2020 10:56:51 GMT
< content-length: 21
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 10001
< server: envoy
<
Finished attempt "1"
* Connection #0 to host 127.0.0.1 left intact
```

But the upstream receives 4 attempts, and the final 3 are cancelled when
the first responds:

```
❯ go run main.go
>>>> attempt 1
>>>> attempt 2
>>>> attempt 3
>>>> attempt 4
slept
<<<< finished 1
request cancelled
<<<< finished 2
request cancelled
<<<< finished 3
request cancelled
<<<< finished 4
```

## how to use

Start the example backend:

```
$ go run main.go
```

Start envoy:

```
# --base-id is there to avoid clashing with the envoy already running in
# your VM
$ envoy -c envoy.yaml --base-id 1 -l debug
```

Make a request to envoy:

```
$ curl --verbose localhost:9091/foo/bar
```
