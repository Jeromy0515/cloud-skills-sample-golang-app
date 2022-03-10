# Cloud Skills Sample GoLang Application
### Language/Framework
##### GoLang/Fiber
##### GoLang Version : 1.17
##### Health Check Path : /health
##### Port : 3000

## How to make go.mod/go.sum files

if you don't have go.mod/go.sum files, install fiber using this command first and import to your application
```
$ go get github.com/gofiber/fiber/v2
```

create go.mod file
```
usage : go mod init [module-name]

$ go mod init
$ ls
go.mod
```

add module requirements and sums
```
usage : go mod tidy [-v]
 
$ go mod tidy 
$ ls 
go.mod go.sum
```

## How to build

you can build using `go build` command
```
usage : go build [file-name]

$ go build app.go
```

## How to test
you can test application using `go test` command


if you use `-v flag` command, you can show all test results  
```
usage: go test [-v flag]

$ go test -v flag
```