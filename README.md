# NIP Checker

## Start HTTP Server
```
$ go run main.go
```

## Start GRPC Server
```
$ go run main.go grpc
```


## Run with Docker
```
$ docker build -t nip-checker .
$ docker run -p 8080:8888 -it nip-checker:latest 
```

## Start Backgound jobs
```
$ go run tasks/tasks.go
```


## Test Rest API
```
$ curl http://localhost:8888/status
```

## Test GRPC Server
```
$ go run grpc/client.go
```