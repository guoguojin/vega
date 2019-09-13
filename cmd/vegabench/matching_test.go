package main

import (
	"testing"
)

func BenchmarkMatching100(b *testing.B) {
	benchmarkMatching(100, b, true, true, "10s")
}

func BenchmarkMatching1000(b *testing.B) {
	benchmarkMatching(1000, b, true, true, "10s")
}

func BenchmarkMatching10000(b *testing.B) {
	benchmarkMatching(10000, b, true, true, "10s")
}

func BenchmarkMatching100000(b *testing.B) {
	benchmarkMatching(100000, b, true, true, "10s")
}

func BenchmarkMatching100Allocated(b *testing.B) {
	benchmarkMatching(100, b, true, true, "10s")
}

func BenchmarkMatching1000Allocated(b *testing.B) {
	b.ReportAllocs()
	benchmarkMatching(1000, b, true, true, "10s")
}

func BenchmarkMatching10000Allocated(b *testing.B) {
	benchmarkMatching(10000, b, true, true, "10s")
}

func BenchmarkMatching100000Allocated(b *testing.B) {
	benchmarkMatching(100000, b, true, true, "10s")
}

func BenchmarkMatching100Uniform(b *testing.B) {
	benchmarkMatching(100, b, true, false, "10s")
}

func BenchmarkMatching1000Uniform(b *testing.B) {
	benchmarkMatching(1000, b, true, false, "10s")
}

func BenchmarkMatching10000Uniform(b *testing.B) {
	benchmarkMatching(10000, b, true, false, "10s")
}

func BenchmarkMatching100000Uniform(b *testing.B) {
	benchmarkMatching(100000, b, true, false, "10s")
}
