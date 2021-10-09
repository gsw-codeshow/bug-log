package main

import (
	proto "bug-log/etcd/proto"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

const schema = "ns"

var (
	ServiceName = flag.String("ServiceName", "greet_service", "service name")        //服务名称
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address") //etcd的地址
)

// etcd解析器
type etcdResolver struct {
	etcdAddr   string
	schema     string
	clientConn resolver.ClientConn
	cli        *clientv3.Client
}

// 初始化一个etcd解析器
func NewResolver(etcdAddr, schema string) resolver.Builder {
	return &etcdResolver{etcdAddr: etcdAddr, schema: schema}
}

func (r *etcdResolver) Scheme() string {
	return r.schema
}

// watch有变化以后会调用
func (r *etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
	fmt.Println(rn)
}

// 解析器关闭时调用
func (r *etcdResolver) Close() {
	log.Println("Close")
}

// 构建解析器 grpc.Dial()同步调用
func (r *etcdResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	// 构建etcd client
	if r.cli == nil {
		r.cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			fmt.Printf("连接etcd失败：%s\n", err)
			return nil, err
		}
	}

	r.clientConn = clientConn
	go r.watch("/" + target.Scheme + "/" + target.Endpoint + "/")

	return r, nil
}

// 监听etcd中某个key前缀的服务地址列表的变化
func (r *etcdResolver) watch(keyPrefix string) {
	// 初始化服务地址列表
	var addrList []resolver.Address

	resp, err := r.cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("获取服务地址列表失败：", err)
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), keyPrefix)})
		}
	}

	r.clientConn.NewAddress(addrList)

	// 监听服务地址列表的变化
	rch := r.cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.clientConn.NewAddress(addrList)
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					r.clientConn.NewAddress(addrList)
				}
			}
		}
	}
}

func exists(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func main() {
	flag.Parse()

	// 注册etcd解析器
	r := NewResolver(*EtcdAddr, schema)
	resolver.Register(r)

	// 客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
	conn, err := grpc.Dial(r.Scheme()+"://author/"+*ServiceName, grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败：", err)
	}
	defer conn.Close()

	// 获得grpc句柄
	c := proto.NewGreetClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for {
			fmt.Println("Morning 调用...")
			// ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
			// defer cancel()
			resp1, err := c.Morning(
				//ctx,
				context.Background(),
				&proto.GreetRequest{Name: "JetWu"},
			)
			if err != nil {
				fmt.Println("Morning调用失败：", err)
				continue
			}
			fmt.Printf("Morning 响应：%s，来自：%s\n", resp1.Message, resp1.From)
			break
		}

		for {
			fmt.Println("Night 调用...")
			resp2, err := c.Night(
				context.Background(),
				&proto.GreetRequest{Name: "JetWu"},
			)
			if err != nil {
				fmt.Println("Night调用失败：", err)
				continue
			}
			fmt.Printf("Night 响应：%s，来自：%s\n", resp2.Message, resp2.From)
			break
		}
	}
}
