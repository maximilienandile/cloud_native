我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Cloud Native Workshop

## Build the application

```
$ go build -o parisApp cmd/paris/main.go
```

With the makefile (recommended)

```
$ make build
```

then launch the binary

```
$ ./bin/paris
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


# Version

Handled via build flags (see `Makefile`)


# Testing

Dockerfile.test can be used to perform checks (linter, test) :

```
$ docker build -test -f Dockerfile.test .
```

# Building a docker image

With a multi-stage docker file :

```
$ docker build -t paris -f Dockerfile .
```
