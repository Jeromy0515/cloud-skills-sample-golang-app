# Cloud Skills Sample GoLang Application
### Language/Framework
##### GoLang/fiber
##### GoLang Version : 1.15
##### Health Check Path : /health
##### Port : 3000
## How to make go.mod/go.sum files
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
if you don't have go.mod/go.sum files, install fiber using this command first
```
$ go get github.com/gofiber/fiber/v2
```

you can build using `go build` command
```
usage : go build [file name]

$ go build app.go
```
