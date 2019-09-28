# Cloud Native Workshop

## Build the application

```
$ go build -o parisApp cmd/paris/main.go
```

## Run the application

```
$ PORT=8080 DIAG_PORT=8081 go run cmd/paris/main.go
```
To check the health of the application

```
$ curl -i http://localhost:8081/health
```

## Configuration

* `PORT` : application port
* `DIAG_PORT` : diagnostic port


# Graceful shutdown

Launch the app :
```
$ PORT=8080 DIAG_PORT=8081 ./parisApp
```

Kill the app with control-c or

```
kill -INT 1234
```

Where 1234 is the process number


# Versioning

