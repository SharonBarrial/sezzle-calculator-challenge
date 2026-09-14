package calculator

import (
	"errors"
	"math"
	"testing"
)

// AlmostEqual checks if two float64 numbers are approximately equal
// considering floating-point precision issues.
func almostEqual(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}

func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
		{"with zero", 0, 5, 5},
		{"decimal numbers", 1.5, 2.25, 3.75},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if !almostEqual(got, tc.expected) {
				t.Errorf("Add(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expected)
			}

		})
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", 3, 5, -2},
		{"with zero", 5, 0, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Subtract(tc.a, tc.b)
			if !almostEqual(got, tc.expected) {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 4, 5, 20},
		{"by zero", 4, 0, 0},
		{"negative numbers", -4, -5, 20},
		{"mixed numbers", -4, 5, -20},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Multiply(tc.a, tc.b)
			if !almostEqual(got, tc.expected) {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !almostEqual(got, 5) {
			t.Errorf("Divide(10, 2) = %v; want 5", got)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if !errors.Is(err, ErrDivisionByZero) {
			t.Errorf("expected ErrDivisionByZero, got %v", err)
		}
	})
}

func TestPower(t *testing.T) {
	cases := []struct {
		name           string
		base, exponent float64
		expected       float64
	}{
		{"positive exponent", 2, 3, 8},
		{"zero exponent", 5, 0, 1},
		{"negative exponent", 2, -2, 0.25},
		{"fractional exponent", 9, 0.5, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Power(tc.base, tc.exponent)
			if !almostEqual(got, tc.expected) {
				t.Errorf("Power(%v, %v) = %v; want %v", tc.base, tc.exponent, got, tc.expected)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	t.Run("perfect square", func(t *testing.T) {
		got, err := Sqrt(9)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !almostEqual(got, 3) {
			t.Errorf("Sqrt(9) = %v; want 3", got)
		}
	})

	t.Run("zero", func(t *testing.T) {
		got, err := Sqrt(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !almostEqual(got, 0) {
			t.Errorf("Sqrt(0) = %v; want 0", got)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		_, err := Sqrt(-4)
		if !errors.Is(err, ErrNegativeSqrt) {
			t.Errorf("expected ErrNegativeSqrt, got %v", err)
		}
	})
}

func TestPercentage(t *testing.T) {
	t.Run("valid percentage", func(t *testing.T) {
		got, err := Percentage(25, 200)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !almostEqual(got, 12.5) {
			t.Errorf("Percentage(25, 200) = %v; want 12.5", got)
		}
	})

	t.Run("zero base", func(t *testing.T) {
		_, err := Percentage(10, 0)
		if !errors.Is(err, ErrInvalidPercentage) {
			t.Errorf("expected ErrInvalidPercentage, got %v", err)
		}
	})

}
