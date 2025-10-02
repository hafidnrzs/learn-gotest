package calculator

import "testing"

func TestAdd(t *testing.T) {
	x, y := 5, 3
	expected := 8
	result := Add(x, y)

	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", x, y, result, expected)
	}
}

func TestSubtract(t *testing.T) {
	x, y := 5, 3
	expected := 2
	result := Subtract(x, y)

	if result != expected {
		t.Errorf("Subtract(%d, %d) = %d; want %d", x, y, result, expected)
	}
}

// Table-driven Tests
func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"zero values", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("got %d; want %d", result, tt.expected)
			}
		})
	}
}

// Helper function
func assertResult(t *testing.T, got, want int, name string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %d; want %d", name, got, want)
	}
}

func TestWithHelper(t *testing.T) {
	assertResult(t, Add(2, 3), 5, "Add(2, 3)")
	assertResult(t, Subtract(5, 3), 2, "Subtract(5, 3)")
}

// Test Complex Types
func TestDivideWithPrecision(t *testing.T) {
	tests := []struct {
		name          string
		a, b          float64
		precision     int
		expected      float64
		expectedError bool
	}{
		{"simple division", 10, 2, 2, 5.00, false},
		{"division with rounding", 10, 3, 3, 3.333, false},
		{"division by zero", 10, 0, 2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &Calculator{precision: tt.precision}
			result, err := calc.DivideWithPrecision(tt.a, tt.b)

			if tt.expectedError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("got %f; want %f", result, tt.expected)
			}
		})
	}
}

// Measure performance with benchmark
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func BenchmarkDivideWithPrecision(b *testing.B) {
	calc := &Calculator{precision: 2}
	for i := 0; i < b.N; i++ {
		calc.DivideWithPrecision(10, 3)
	}
}
