package main

import (
	"testing"
)

func BenchmarkWithMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usingMake()
	}
}

func BenchmarkWithVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usingVar()
	}
}
