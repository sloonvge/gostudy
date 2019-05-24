package main

import (
	"log"
)

type BytesPool struct {
	pool chan []byte
}

func NewBytesPool(max int) *BytesPool {
	return &BytesPool{
		pool: make(chan []byte, max),
	}
}

func (bp *BytesPool) Get(s int) []byte {
	var c []byte

	select {
	case c = <- bp.pool:
	default:
		log.Println("new []byte")
		return make([]byte, s)
	}

	if cap(c) < s {
		return make([]byte, s)
	}

	return c[:s]
}

func (bp *BytesPool) Put(c []byte) {
	select {
	case bp.pool <- c:
		log.Println("input")
	default:


	}
}

func main() {
	pool := NewBytesPool(2)
	buf := []struct{
		b []byte
	} {
		{b: []byte("abc")},
		{b: []byte("123D")},
		{b: []byte("ABC")},
	}

	for _, b := range buf {
		pool.Put(b.b)
	}

	for i := 0; i< len(buf); i++ {
		log.Println(string(pool.Get(len(buf[i].b))))
	}
}
