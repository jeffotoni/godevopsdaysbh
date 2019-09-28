package main

import "testing"

var str string

func BenchmarkStringBuildNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = Native()
	}
}
func BenchmarkStringBuildBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = Builder()
	}
}
