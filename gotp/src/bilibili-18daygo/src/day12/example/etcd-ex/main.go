package main

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed err:", err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s: %s\n", kv.Key, kv.Value)
	}
}
