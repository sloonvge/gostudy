package main

import (
	"fmt"
	etcdClient "github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect etcd err:%v\n", err)
		return
	}

	fmt.Println("connect etcd success")
	defer cli.Close()
}
