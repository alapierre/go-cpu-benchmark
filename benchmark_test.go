package main

import (
	"crypto/rand"
	"crypto/sha256"
	"sync"
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
	benchmarkSha(Kilo, 1000, b)
}

func BenchmarkSha1m(b *testing.B) {
	benchmarkSha(Mega, 1000, b)
}

func benchmarkSha(len, amount int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiTh(amount, len)
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

func multiTh(amount, len int) {

	var wg sync.WaitGroup
	wg.Add(amount)

	for i := 0; i < amount; i++ {
		go func() {
			defer wg.Done()
			makeSha(len)
		}()
	}
	wg.Wait()
}
