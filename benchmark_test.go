package main

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

const (
	Kilo int = 1024
	Mega     = Kilo * 1024
	Giga     = Mega * 1024
	Tera     = Giga * 1024
	Peta     = Tera * 1024
	Exa      = Peta * 1024
)

func BenchmarkSha1k(b *testing.B) {
	benchmarkSha(Kilo, b)
}

func BenchmarkSha1m(b *testing.B) {
	benchmarkSha(Mega, b)
}

func BenchmarkSha1g(b *testing.B) {
	benchmarkSha(Giga, b)
}

func benchmarkSha(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSha(len)
	}
}

func makeSha(len int) {
	source := make([]byte, len)
	_, err := rand.Read(source)
	if err != nil {
		panic("cant read random")
	}

	h := sha256.New()
	h.Write(source)
}