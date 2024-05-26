package utesting

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestGetMax(t *testing.T) {
	table := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{5, 4, 5},
		{3, 3, 3},
	}
	for _, tt := range table {
		result := GetMax(tt.a, tt.b)
		if result != tt.want {
			t.Errorf("GetMax(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		n, want int
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
	}
	for _, tt := range table {
		result := Fibonacci(tt.n)
		if result != tt.want {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, result, tt.want)
		}
	}
}
