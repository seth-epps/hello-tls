# hello-go
Simple json webserver written in go with [gorilla/mux](https://github.com/gorilla/mux)


## Usage
```sh
docker image build -t hello-go .
docker container run --rm -d -p 8080:8080 --name hello-go hello-go
curl localhost:8080/
# {"ip":"172.17.0.1:55532","message":"Hello From Go!"}
```