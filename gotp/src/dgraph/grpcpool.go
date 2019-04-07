package dgraph

import (
	"google.golang.org/grpc"
	"log"
)

type GrpcPoll struct {
	pool chan *grpc.ClientConn
}

func NewGrpcPoll(n int) *GrpcPoll {
	return &GrpcPoll{
		pool: make(chan *grpc.ClientConn, n),
	}
}

func (gp *GrpcPoll) Put(gc *grpc.ClientConn) {
	select {
	case gp.pool <- gc:
		log.Println("add to pool")
	default:

	}
}

func (gp *GrpcPoll) Get() (gc *grpc.ClientConn) {
	select {
	case gc = <- gp.pool:
		log.Println("get one conn")
	default:
		log.Println("empty pool")
	}

	return
}
