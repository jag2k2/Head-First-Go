# how build build and run examples in this directory:

go version
go fmt ./hello/main.go
go run ./hello/main.go
go build -o ./hello/bin/hello ./hello/main.go 
./hello/bin/hello
