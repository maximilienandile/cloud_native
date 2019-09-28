# Cloud Native Workshop

## Run the application

```
$ PORT=8080 DIAG_PORT=8081 go run cmd/paris/main.go
```
To check the health of the application

```
$ curl -i http://localhost:8081/health
```

## Configuration

* PORT : application port
* DIAG_PORT : diagnostic port
