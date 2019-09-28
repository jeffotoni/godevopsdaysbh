package main

import "testing"
import "sync"
import "fmt"

var m sync.Map
var ch = make(chan string)

func BenchmarkProduceMap(b *testing.B) {
	for N := 0; N < b.N; N++ {
		for i := 0; i < 10000000; i++ {
			m.LoadOrStore(N, string(N))
		}
	}
}

func BenchmarkConsumeMap(b *testing.B) {
	for N := 0; N < b.N; N++ {
		for i := 0; i < 10000000; i++ {
			v, ok := m.LoadOrStore(N, string(N))
			if ok {
				fmt.Sprintf("Consumer: %s", v)
			}
		}
	}
}

func BenchmarkProduceChan(b *testing.B) {
	for N := 0; N < b.N; N++ {
		go func() {
			for i := 0; i < 10000000; i++ {
				ch <- string(i)
			}
		}()
	}
}

func BenchmarkConsumeChan(b *testing.B) {
	go func() {
		for v := range ch {
			fmt.Sprintf("Consumer: %s", v)
		}
	}()
}

// func BenchmarkConsumeChan(b *testing.B) {
// 	go func() {
// 		for {
// 			select {
// 			case v := <-ch:
// 				fmt.Sprintf("Consumer: %s", v)
// 			}
// 		}
// 	}()
// }
