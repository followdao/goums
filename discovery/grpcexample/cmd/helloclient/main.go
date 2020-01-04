package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/integrii/flaggy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"github.com/tsingson/goums/discovery/etcdv3lb"
	"github.com/tsingson/goums/discovery/grpcexample/helloworld"
)

func main() {

	srv := "hello_service"
	addr := "http://127.0.0.1:2379"
	flaggy.String(&srv, "s", "service", "")
	flaggy.String(&addr, "a", "address", "")
	flaggy.Parse()

	fmt.Println("server: ", srv, "  address: ", addr)

	// r := etcdlb.NewResolver(srv)
	//
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// conn, err := grpc.DialContext(ctx, addr,
	// 	grpc.WithBalancer(grpc.RoundRobin(r)),
	// 	grpc.WithBlock(), // 客户端将连接到 GPRC 服务, 直到连接成功
	// 	grpc.WithInsecure(),
	// 	// grpc.WithCodec(flat.FlatbuffersCodec{}),
	// 	// grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{})),
	// 	// client side
	// 	grpc.WithInitialWindowSize(grpcInitialWindowSize),
	// 	grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
	// 	grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
	// 	grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
	// 	grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: grpcBackoffMaxDelay}),
	// 	grpc.WithKeepaliveParams(keepalive.ClientParameters{
	// 		Time:                grpcKeepAliveTime,
	// 		Timeout:             grpcKeepAliveTimeout,
	// 		PermitWithoutStream: true,
	// 	}))
	// if err != nil {
	// 	panic(err)
	// }

	// 连接etcd,得到名命名空间
	schema, err := etcdv3lb.GenerateAndRegisterEtcdResolver("127.0.0.1:2379", "HelloService")
	if err != nil {
		//	log.Fatal("init etcd resolver err:", err.Error())
		os.Exit(-1)
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:///HelloService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		fmt.Println(err)
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		client := helloworld.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
		}
	}
}
