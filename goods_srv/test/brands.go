package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err := grpc.Dial("192.168.50.210:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)

}

func TestGetBrandList() {
	resp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       2,
		PagePerNums: 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Total)

	for _, brand := range resp.Data {
		fmt.Println(brand.Name)
	}
}

func main() {
	Init()
	TestGetBrandList()
	conn.Close()
}
