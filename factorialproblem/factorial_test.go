package factorialproblem

import "testing"

func TestReturnFactorial(t *testing.T) {
	want := 120
	actual := factorial(5)
	if want != actual {
		t.Errorf("want %d, actual %d", want, actual)
	}
}

func TestReturnFactorialOf1(t *testing.T) {
	want := 1
	actual := factorial(1)
	if want != actual {
		t.Errorf("want %d, actual %d", want, actual)
	}
}

func TestReturnFactorialOf0(t *testing.T) {
	want := 1
	actual := factorial(0)
	if want != actual {
		t.Errorf("want %d, actual %d", want, actual)
	}
}

func TestReturnFactorialOfNegativeNumber(t *testing.T) {
	want := 0
	actual := factorial(-1)
	if want != actual {
		t.Errorf("want %d, actual %d", want, actual)
	}
}

func TestReturnFactorialOfAnyNumber(t *testing.T) {
	tests := []struct {
		Description string
		want        int
		actual      int
	}{
		{"Factorial of 5 is 120", 5, 120},
		{"Factorial of 1 is 1", 1, 1},
		{"Factorial of 0 is 1", 0, 1},
		{"Factorial of -1 is 0", -1, 0},
	}

	for _, test := range tests {
		if test.actual != factorial(test.want) {
			t.Errorf(test.Description)
		}
	}
}
