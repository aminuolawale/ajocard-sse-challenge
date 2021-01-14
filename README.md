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
$ docker build -t aminuolawale/nip-checker .
$ docker run -p 8080:8888 -it nip-checker:latest 
```