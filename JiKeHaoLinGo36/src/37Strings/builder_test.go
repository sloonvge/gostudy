package main

import "testing"

func BenchmarkE_Join(b *testing.B) {
	var t E
	t = "2.05"
	s := "Kg"

	for i := 0; i < b.N; i++ {
		t.Join(s)
	}
}

func BenchmarkE_Builder(b *testing.B) {
	var t E
	t = "2.05"
	s := "Kg"

	for i := 0; i < b.N; i++ {
		t.Join(s)
	}
}
