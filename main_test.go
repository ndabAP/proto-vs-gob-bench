package main

import "testing"

func BenchmarkGobWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gobWrite()
	}
}

func BenchmarkGobRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gobRead()
	}
}

func BenchmarkProtobufWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		protoWrite()
	}
}

func BenchmarkProtobufRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		protoRead()
	}
}
