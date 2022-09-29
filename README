# ipecho

A very basic webserver that listens on port 80 and returns the apparent origin IP.

Built with [gin](https://github.com/gin-gonic/gin).

## Step 1: Build

Compiling requires [Golang](https://go.dev/doc/install).

Clone the repo, move into the directory and run:

```
go build
```

This will compile to the binary file `ipecho`.

## Step 2: Start listening

Port 80 requires root permissions, so you need to run the binary as root. For example:

```
sudo ./ipecho
```

Logging is done to STDOUT.

## Step 3: Examples

```
$ curl localhost
HTTP Error 400: Bad Request.
Supported endpoints are /json and /text
```

```
$ curl -s localhost/json | jq
{
  "IP": "127.0.0.1",
  "USERAGENT": "curl/7.85.0"
}
```

```
$ curl -s localhost/text
127.0.0.1
```