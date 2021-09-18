# Redirect

Simple HTTP server that sends redirect responses. Useful when leveraging **SSRF** attacks.



## Usage

### Redirect customization

**Specify host:** At least you have to specify the host to redirect to with a positonal argument:
```shell
redirect http://127.0.0.1
```

This will redirect to 127.0.0.1 with the same path requested.

**Specify schema:** If no options are provided the redirect will default to HTTP. If you want to redirect to HTTPS, FTP, Gopher or anything else you can just add the schema to the host:

```shell
redirect ftp://127.0.0.1:2121
```

**Specifically set a path:** Maybe you just want to set a path. The `-path` flag sets a path to redirect:

```shell
redirect -path /custom/path  ftp://127.0.0.1:2121
```

**Custom status code:** By default redirect responses are served with status 301. Change it with the `-status` flag:

```shell
redirect -status 302 https://127.0.0.1:8443
```


### Server configuration

**Change listening address and port:** By default the server listens at `0.0.0.0:8888`, but you can change it with the `-addr` flag:

```shell
redirect -addr :8080 https://example.com
```


```shell
redirect -addr 127.0.0.1:8080 https://example.com
```

**Dump the requests to stdout:** The `-dump` flag lets you see what requests are reaching the server

```shell
redirect -dump ftp://127.0.0.1:25
```


**Dump requests to a file:** Just use the `-dump` flag with shell redirection `>` or the `tee` command:


```shell
redirect -dump http://127.0.0.1:25 > requests.log
redirect -dump http://127.0.0.1:25 | tee requests.log
```

## Installation

With the newest `go install`:
```shell
go install github.com/luastan/redirect@latest
```

With good old `go get`:

```shell
go get github.com/luastan/redirect
```

## TODOs

- [ ] Support TLS on the listener