package main

import (
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"google.golang.org/grpc"
	"log"
	"net"
	dal "github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

func main() {
	dal.InitMongoDB()
	// todo constants and others
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8889"))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	biz.RegisterBizServiceServer(srv, &BizServerImpl{})
	r, err := etcd.NewEtcdRegistry([]string{"123.56.152.169:2379"})
	err = r.Register(&registry.Info{
		ServiceName: "tiktok_biz",
		Addr: &net.TCPAddr{
			IP:   []byte{127, 0, 0, 1},
			Port: 8889,
		},
	})
	if err != nil {
		log.Panic(err)
	}
	srv.Serve(lis)
}
