package dgraph

import (
	"testing"
	"fmt"
	"context"
	"sync"
)

func TestConn(t *testing.T) {

	var wg sync.WaitGroup
	n := 10
	wg.Add(n)


	for i := 0; i < n; i++ {

		go func() {
			dg, cancel := getDgraphClient(DBURL)
			defer cancel()
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
			fmt.Println(string(resp.Json))

			wg.Done()
		}()
	}

	wg.Wait()

}
