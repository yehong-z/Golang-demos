# Grp简单demo

基本调用，批量调用，流调用

grpc基于protocol， 需要先安装相应的编译程序
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# 如果没有加入环境变量需要额外操作
export PATH="$PATH:$(go env GOPATH)/bin"
```

编译proto文件
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  protoc/test.proto
```

```shell
go run ./server/main.go
```

```shell
go run ./client/main.go
```
