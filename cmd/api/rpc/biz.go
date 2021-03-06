package rpc

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	etcd "github.com/kitex-contrib/registry-etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var BizClient biz.BizServiceClient

func initBiz() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	resolver, err := etcd.NewEtcdResolver([]string{"123.56.152.169:2379"})
	if err != nil {
		log.Panic(err)
	}
	resolve, err := resolver.Resolve(context.Background(), "tiktok_biz")
	if err != nil {
		return
	}
	if len(resolve.Instances) == 0 {
		return
	}
	addr := resolve.Instances[0].Address().String()
	//conn, err := grpc.Dial("127.0.0.1:8888", opts...)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
	}
	BizClient = biz.NewBizServiceClient(conn)
}
