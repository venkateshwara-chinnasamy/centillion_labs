package main

import (
	"context"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	// Successful processing test
	ctx := context.Background()
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, err := Pipeline(ctx, input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 1 + 4 + 9 + 16 + 25 + 36 + 49 + 64 + 81 + 100
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Context cancellation test
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	input = []int{10, 20, 30, 40, 50}
	_, err = Pipeline(ctx, input)
	if err != context.DeadlineExceeded {
		t.Errorf("Expected context DeadlineExceeded, got %v", err)
	}
}

func BenchmarkPipeline(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Pipeline(ctx, input)
		if err != nil {
			b.Fatalf("Benchmark error: %v", err)
		}
	}
}
