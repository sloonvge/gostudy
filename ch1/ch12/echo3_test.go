package ch12

import "testing"

func BenchmarkEchoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEchoStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
