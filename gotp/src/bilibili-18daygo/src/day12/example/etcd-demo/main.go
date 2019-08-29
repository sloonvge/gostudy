package main

import (
	"fmt"
	etcdClient "github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"localhost:3237"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connected etcd err:%v\n", err)
	}
	fmt.Println("connected etcd success")
	defer cli.Close()
}
