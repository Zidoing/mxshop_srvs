package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/initialize"
	"mxshop_srvs/user_srv/proto"
	"mxshop_srvs/user_srv/utils"
	"net"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()

	IP := flag.String("ip", "0.0.0.0", "ip  地址")
	PORT := flag.Int("port", 0, "端口号")
	flag.Parse()

	zap.S().Info("ip", *IP, "port:", *PORT)

	if *PORT == 0 {
		*PORT, _ = utils.GetFreePort()
	}

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *PORT))

	if err != nil {
		panic("failed to listen" + err.Error())
	}

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
	)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "172.100.22.12", *PORT),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *PORT
	registration.Tags = []string{"zhoulei", "go"}
	registration.Address = "172.100.22.12"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listener)

	if err != nil {
		panic("failed to start grpc " + err.Error())
	}
}
