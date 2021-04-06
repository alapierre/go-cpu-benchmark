package main

import (
	"crypto/rand"
	"crypto/sha256"
	"sync"
	"testing"
)

//goland:noinspection ALL
const (
	Kilo int = 1024
	Mega     = Kilo * 1024
	Giga     = Mega * 1024
	Tera     = Giga * 1024
	Peta     = Tera * 1024
	Exa      = Peta * 1024
)

func BenchmarkParallelSha1k(b *testing.B) {
	benchmarkSha(Kilo, 1000, b)
}

func BenchmarkParallelSha1m(b *testing.B) {
	benchmarkSha(Mega, 1000, b)
}

func BenchmarkSeqSha1m(b *testing.B) {
	benchmarkSeqSha(Mega, b)
}

func benchmarkSha(len, amount int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiTh(amount, len)
	}
}

func benchmarkSeqSha(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSha(len)
	}
}

func BenchmarkRunParallelSha1m(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			makeSha(Mega)
		}
	})
}

func makeSha(len int) {

	//fmt.Printf("\rthreads: %d", runtime.NumGoroutine())

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
