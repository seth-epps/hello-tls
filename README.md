# hello-tls

Simple json webserver written in go

## Usage

### Certificate setup

Generate self-signed server certificate

```sh
❯ ./cert-gen.sh hello.tls "*.hello.tls"
Certificate request self-signature ok
subject=CN=hello.tls
```

### Run

#### Local Build

```sh
go build -o .dist/serve
.dist/serve -tls-key .certs/server.key -tls-cert .certs/server.crt -port 7443
2024/03/19 21:54:29 Server starting on :7443

## In another shell
curl -k https://localhost:7443/
{"ip":"[::1]:62827","message":"Hello From Go w/ TLS!"}
```

#### Docker

```sh
docker image build -t hello-tls .
docker container run --rm -d \
    -p 7443:8443 \
    --mount type=bind,source="$(pwd)"/.certs,target=/etc/certs \
    --name hello-tls \
    hello-tls \
    -tls-key /etc/certs/server.key \
    -tls-cert /etc/certs/server.crt
a837f07e6a6978206c70afd7ee57b043dbbf030788b4963c545f0ffa0c370b5f

docker container ls -a
CONTAINER ID   IMAGE                  COMMAND                  CREATED         STATUS         PORTS                                                                       NAMES
a837f07e6a69   hello-tls              "/app -tls-key /etc/…"   6 seconds ago   Up 5 seconds   0.0.0.0:7443->8443/tcp                                                      hello-tls

curl -k https://localhost:7443/
{"ip":"192.168.65.1:37881","message":"Hello From Go w/ TLS!"}
```
