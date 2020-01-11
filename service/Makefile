build:
# 生成go文件
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/asifs/service proto/filter/filter.proto
# 生成php文件

# 根据当前目录下的 Dockerfile 生成名为 asifs 的镜像
#	GOOS=linux GOARCH=amd64 go build
#	docker build -t asifs .

http:
	go run http_server.go

rpc:
	go run rpc_server.go


