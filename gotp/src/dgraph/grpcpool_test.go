package dgraph

import (
	"testing"
	"google.golang.org/grpc"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgo"
	"golang.org/x/net/context"
	"log"
	"time"
	"sync"
)

func TestNewGrpcPoll(t *testing.T) {
	pools := NewGrpcPoll(3)




	var query = func() {
		conn, err := grpc.Dial(DBURL, grpc.WithInsecure())
		if err != nil {
			t.Log(err)
		}

		pools.Put(conn)
		c := pools.Get()
		defer c.Close()
		dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
		after := time.Second * 2

		ticker := time.NewTicker(after)
		defer ticker.Stop()

		for {
			<- ticker.C
			ctx := context.Background()
			txn := dg.NewTxn()
			defer txn.Discard(ctx)

			q := `
		{
		  me(func: eq(node_type, 9)) {
		fsi:count(uid)
		  }
		}
	`
			resp, err := txn.Query(ctx, q)
			if err != nil {
				t.Log(err)
			}
			log.Println(string(resp.Json))
		}

	}

	var wg sync.WaitGroup
	n := 4
	wg.Add(n)
	for i:=0; i < n; i++{
		go query()
	}

	wg.Wait()

}
