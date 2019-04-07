package dgraph

import (
	"github.com/dgraph-io/dgo"
	"google.golang.org/grpc"
	"log"
	"github.com/dgraph-io/dgo/protos/api"
	"context"
)

// sync.pool

func getDgraphClient(target string) (*dgo.Dgraph, context.CancelFunc) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	dc := api.NewDgraphClient(conn)
	return dgo.NewDgraphClient(dc), func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error: %v", err)
		}
	}
}

//func getDgraphClient(target string) (*dgo.Dgraph, *grpc.ClientConn) {
//	conn, err := grpc.Dial(target, grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	dc := api.NewDgraphClient(conn)
//	return dgo.NewDgraphClient(dc), conn
//}

