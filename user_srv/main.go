package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/proto"
	"net"
)

func main() {

	IP := flag.String("ip", "0.0.0.0", "ip  地址")
	PORT := flag.Int("port", 50051, "端口号")
	flag.Parse()

	fmt.Println("ip", *IP, "port:", *PORT)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *PORT))

	if err != nil {
		panic("failed to listen" + err.Error())
	}

	err = server.Serve(listener)

	if err != nil {
		panic("failed to start grpc " + err.Error())
	}
}
