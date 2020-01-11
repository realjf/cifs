package main

import (
	pb "asifs/service/proto/filter"
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	ADDRESS           = "localhost:40001"
)

func main() {

	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()

	// 初始化gRPC客户端
	client := pb.NewFilterServiceClient(conn)

	request := &pb.Request{
		Content: "hello, 你好，我是，fuck you",
	}

	// 过滤停用词
	response, err := client.StopWord(context.Background(), request)
	if err != nil {
		log.Printf("stop word filter error: %v", err)
	}

	log.Printf("filtered_content: %s", response.Result.FilteredContent)

	// 过滤敏感词
	response, err = client.SensitiveWord(context.Background(), request)
	if err != nil {
		log.Printf("sensitive word filter error: %v", err)
	}

	log.Printf("filtered_content: %s", response.Result.FilteredContent)
}


