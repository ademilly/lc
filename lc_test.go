package main

import "testing"
import "strings"

// generateBigString generates a big string made of 10k "hello" joined by "\n"
// returns "hello\n" x 10k
func generateBigString() string {
	arr := make([]string, 10000)

	for i := range arr {
		arr[i] = "hello"
	}

	return strings.Join(arr, "\n")
}

// TestCountLines tests countLines function
func TestCountLines(t *testing.T) {
	n := countLines(strings.NewReader("hello\nworld\n"))
	if n != 2 {
		t.Error("expected 2, got:", n)
	}
}

// TestCountLinesBig tests countLines using the generateBigString function
// simulates countLines usage with a file of 10k lines
func TestCountLinesBig(t *testing.T) {
	n := countLines(strings.NewReader(generateBigString()))
	if n != 10000 {
		t.Error("expected 2, got:", n)
	}
}

// BenchmarkCountLines benchmarks countLines using generateBigString function
// it mostly exists for shits and giggles because benchmarking and testing is fun in go
func BenchmarkCountLines(b *testing.B) {
	setup := generateBigString()
	for i := 0; i < b.N; i++ {
		_ = countLines(strings.NewReader(setup))
	}
}
