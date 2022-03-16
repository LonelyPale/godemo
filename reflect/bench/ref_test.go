package main

import "testing"

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkNewUseReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewUseReflect()
	}
}

func BenchmarkUseValue(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UseValue()
	}
}
