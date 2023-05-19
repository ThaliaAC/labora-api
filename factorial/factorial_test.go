package factorial

import (
	"fmt"
	"testing"
)

type factorialTest struct {
	arg1, expected int64
}

var factorialTests = []factorialTest{
	{0, 1},
	{4, 24},
	{6, 720},
}

func TestFactorial(t *testing.T) {
	for _, test := range factorialTests {
		if output := Factorial(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(3)
	}
}

func ExampleFactorial() {
	fmt.Println(Factorial(3))
	// Output: 6
}
