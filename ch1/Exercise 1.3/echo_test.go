package main

import "testing"

func BenchmarkEchoSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoSlow()
	}
}

func BenchmarkEchoFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoFast()
	}
}
