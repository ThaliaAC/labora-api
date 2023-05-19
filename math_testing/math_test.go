package math_operations

import (
	"fmt"
	"testing"
)

/*func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}*/

// arg1 significa el argumento 1 and arg2 el argumento 2, and the expected stands for the 'el resultado que esperamos'
type addTest struct {
	arg1, arg2, expected int
}

type subtractTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

var subtractTests = []subtractTest{
	{3, 2, 1},
	{8, 4, 4},
	{9, 3, 6},
	{10, 3, 7},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}

func TestSubtract(t *testing.T) {
	for _, test := range subtractTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(15, 10)
	}
}

func ExampleSubtract() {
	fmt.Println(Subtract(15, 10))
	// Output: 5
}
