# GRPC_DEMO
用gRPC框架写一个mysql调用的demo


## 1 执行步骤

### 安装protoc

1. 首先安装protobuf

``` bash
brew install protobuf
```

2. 安装支持go protoc的插件

> protoc原生是不支持go语言的自动代码生成，需要手动下载一个插件，然后在命令里面输入--go_out关键字时会执行插件，生成go的代码。

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26  #好像不好用，换下面的
go get github.com/golang/protobuf/protoc-gen-go 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

3. 安装插件后，更新环境变量

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 写proto文件，生成代码

1. 这里有一点要注意，在代码里面要加上生成文件的路径。（按道理来说，执行protoc命令的时候带上输出路径了，这里不加的话会报错）。

```protobuf
option go_package = "/gen";
```

2. 执行生成代码的命令

```bash
protoc -I proto toupper.proto --go_out=plugins=grpc:.
```

### build代码，并运行

```bash
make
./server
./client #另起一个终端
```



