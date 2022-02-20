# healthcheck-go

## run

```sh
$ docker build -t healthcheck-go .
$ docker run -p 80:80 --rm healthcheck-go
```

## push

```sh
$ docker tag healthcheck-go:latest ${REPOSITORY}:latest
$ docker push ${REPOSITORY}:latest
```
