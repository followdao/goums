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

	// ticker := time.NewTicker(time.Duration(500) * time.Millisecond)
	// for t := range ticker.C {
	client := helloworld.NewGreeterClient(conn)
	for {
		resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "world " + strconv.FormatInt(time.Now().Unix(), 10)})
		if err == nil {
			fmt.Printf(" Reply is %s\n", resp.Message)
		}
	}
	// }
}
