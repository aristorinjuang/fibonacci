package fibonacci

import (
	"testing"
)

var samples = []struct {
	index uint
	value uint
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{11, 89},
	{12, 144},
	{15, 610},
	{20, 6765},
	{30, 832040},
	{45, 1134903170},
}

func TestRecursion(t *testing.T) {
	for _, sample := range samples {
		value := Recursion(sample.index)

		if value != sample.value {
			t.Errorf("expected %d, got %d", sample.value, value)
		}
	}
}

func BenchmarkRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range samples {
			Recursion(sample.index)
		}
	}
}

func TestMemoization(t *testing.T) {
	for _, sample := range samples {
		value := Memoization(sample.index)

		if value != sample.value {
			t.Errorf("expected %d, got %d", sample.value, value)
		}
	}
}

func BenchmarkMemoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range samples {
			Memoization(sample.index)
		}
	}
}

func TestIteration(t *testing.T) {
	for _, sample := range samples {
		value := Iteration(sample.index)

		if value != sample.value {
			t.Errorf("expected %d, got %d", sample.value, value)
		}
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range samples {
			Iteration(sample.index)
		}
	}
}

func TestMatrix(t *testing.T) {
	for _, sample := range samples {
		value := Matrix(sample.index)

		if value != sample.value {
			t.Errorf("expected %d, got %d", sample.value, value)
		}
	}
}

func BenchmarkMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range samples {
			Matrix(sample.index)
		}
	}
}

func TestMath(t *testing.T) {
	for _, sample := range samples {
		value := Math(sample.index)

		if value != sample.value {
			t.Errorf("expected %d, got %d", sample.value, value)
		}
	}
}

func BenchmarkMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range samples {
			Math(sample.index)
		}
	}
}
