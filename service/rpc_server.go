package main

import (
	"asifs/service/config"
	"asifs/service/db"
	"asifs/service/filters"
	pb "asifs/service/proto/filter"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

const (
	port = ":40001"
)


var (
	Config *config.Config
	pprof_address = "0.0.0.0:6060"
)

func init() {
	// config、db
	Config = config.NewConfig().LoadConfig("./config/config.json")
	db.NewMysql(Config).Init()
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		if err := http.ListenAndServe(pprof_address, nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterFilterServiceServer(server, filters.NewFilter())

	log.Println("listen on: ", port)
	log.Fatal(server.Serve(listen))
}
