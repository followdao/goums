package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/integrii/flaggy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tsingson/goums/discovery/etcdv3lb"
	"github.com/tsingson/goums/discovery/grpcexample/helloworld"
)

func main() {

	srv := "hello_service"
	port := 50001
	addr := "http://127.0.0.1:2379"

	flaggy.String(&srv, "s", "service", "")
	flaggy.String(&addr, "a", "address", "")
	flaggy.Int(&port, "p", "port", "")

	flaggy.Parse()

	/**
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	err = etcdlb.Register(srv, "127.0.0.1", port, addr, time.Second*10, 15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		etcdlb.UnRegister()
		os.Exit(1)
	}()


	*/

	lis, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//构造注册中心对象
	etcdRegister := etcdv3lb.NewRegister("127.0.0.1:2379")

	//开始注册
	go func() {
		for {
			etcdRegister.Register(etcdv3lb.ServiceMetadata{ServiceName: "HelloService",
				Host: "127.0.0.1", Port: port, IntervalTime: time.Duration(10)})

			time.Sleep(time.Second * 5)
		}
	}()

	log.Printf("starting hello service at %d", port)
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}
