package main

import (
	proto "bug-log/etcd/proto"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const schema = "ns"

var (
	Port        = flag.Int("Port", 3000, "listening port")                           //服务器监听端口
	ServiceName = flag.String("ServiceName", "greet_service", "service name")        //服务名称
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address") //etcd的地址
	Host        = flag.String("Host", "127.0.0.1", "host")
)

type Server struct {
	cli         *clientv3.Client //etcd client
	Schema      string
	EtcdAddr    string
	ServiceName string
	ServerAddr  string
}

//将服务地址注册到etcd中
func (s *Server) Register(ttl int64) (err error) {
	if nil == s.cli {
		s.cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(s.EtcdAddr, ";"),
			DialTimeout: 60 * time.Second,
		})
		if nil != err {
			// 连接etcd失败
			fmt.Println("连接etcd失败: %s\n", err)
			return
		}
	}
	// 与etcd建立长连接，并保证连接不断
	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(ttl))
		key := "/" + s.Schema + "/" + s.ServiceName + "/" + s.ServerAddr
		for {
			resp, err := s.cli.Get(context.Background(), key)
			if nil != err {
				fmt.Println("获取服务器地址失败：%s", err)
			} else if 0 == resp.Count {
				err = s.keepAlive(ttl)
				if nil != err {
					fmt.Println("保持连接失败：%s", err)
				}
			}
			_, ok := <-ticker.C
			if !ok {
				break
			}
		}
	}()

	//关闭信号处理
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		sign := <-ch
		s.UnRegister()
		if i, ok := sign.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	return
}

//保持服务器与etcd的长连接
func (s *Server) keepAlive(ttl int64) error {
	// 创建租约
	leaseResp, err := s.cli.Grant(context.Background(), ttl)
	if nil != err {
		fmt.Printf("创建租期失败：%s\n", err)
		return err
	}

	// 将服务地址注册到etcd中
	key := "/" + s.Schema + "/" + s.ServiceName + "/" + s.ServerAddr
	_, err = s.cli.Put(context.Background(), key, s.ServerAddr, clientv3.WithLease(leaseResp.ID))
	if nil != err {
		fmt.Printf("注册服务失败：%s", err)
		return err
	}

	// 建立长连接
	responChan, err := s.cli.KeepAlive(context.Background(), leaseResp.ID)
	if nil != err {
		fmt.Printf("建立长连接失败：%s\n", err)
		return err
	}

	go func() {
		// 清空keepAlive返回的channel
		for {
			_, ok := <-responChan
			if !ok {
				break
			}
		}
	}()
	return nil
}

//取消注册
func (s *Server) UnRegister() {
	if nil != s.cli {
		key := "/" + s.Schema + "/" + s.ServiceName + "/" + s.ServerAddr
		s.cli.Delete(context.Background(), key)
	}
	return
}

//rpc服务接口
type greetServer struct{}

func (gs *greetServer) Morning(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Printf("Morning 调用: %s\n", req.Name)
	time.Sleep(time.Second * (time.Duration)(3*(rand.Int()%10+1)))
	fmt.Println("---moring---")
	return &proto.GreetResponse{
		Message: "Good morning, " + req.Name,
		From:    fmt.Sprintf("%s:%d", *Host, *Port),
	}, nil
}

func (gs *greetServer) Night(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Printf("Night 调用: %s\n", req.Name)
	time.Sleep(time.Second * (time.Duration)(3*(rand.Int()%10+1)))
	return &proto.GreetResponse{
		Message: "Good night, " + req.Name,
		From:    fmt.Sprintf("%s:%d", *Host, *Port),
	}, nil
}

func main() {
	flag.Parse()

	//将服务地址注册到etcd中
	serverAddr := fmt.Sprintf("%s:%d", *Host, *Port)

	//监听网络
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return
	}
	defer listener.Close()

	//创建grpc句柄
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//将greetServer结构体注册到grpc服务中
	proto.RegisterGreetServer(srv, &greetServer{})
	fmt.Printf("greeting server address: %s\n", serverAddr)

	s := Server{
		EtcdAddr:    *EtcdAddr,
		ServiceName: *ServiceName,
		ServerAddr:  serverAddr,
		Schema:      schema,
	}
	s.Register(5)

	//监听服务
	err = srv.Serve(listener)
	if err != nil {
		fmt.Println("监听异常：", err)
		return
	}
}
