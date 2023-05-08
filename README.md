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
$ curl -s localhost/json | jq
{
  "IP": "127.0.0.1",
  "USERAGENT": "curl/7.85.0"
}
```

```
$ curl -s localhost
127.0.0.1
```

## Dockerfile
In the same directory as `Dockerfile`, run:
```
$ docker build --tag ipecho .
(...)
Successfully built 2eeb58f23dac
Successfully tagged ipecho:latest
```
```
$ docker image ls
REPOSITORY   TAG               IMAGE ID       CREATED          SIZE
ipecho       latest            2eeb58f23dac   14 minutes ago   10.4MB
```
Start a container from the image (`latest` is implicit):
```
$ docker run --network=bridge -d ipecho
d240f5042683f299915b8c16f5c45cf2642ab4d69385c4c0c92a2523929ded1c
```
Figure out the IP and test it:
```
$ docker inspect d2 | grep IP
(...)
$ curl 172.17.0.2
172.17.0.1
```
