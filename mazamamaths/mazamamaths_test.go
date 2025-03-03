package mazamamaths_test

import (
	"pacx/mazamamaths"
	"testing"
)

type addTestCase struct {
	a, b, expected int
}

var testCases = []addTestCase{
	{1, 2, 3},
	{2, 2, 4},
	{25, 25, 50},
	{1, 24, 25},
}

func TestAdd(t *testing.T) {

	//got := mazamamaths.Add(1, 2)
	//expected := 2
	//
	//if got != expected {
	//	t.Fail()
	//}

	for _, tc := range testCases {
		got := mazamamaths.Add(tc.a, tc.b)

		if got != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, got)
		}
	}

}

func FuzzTestAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		mazamamaths.Add(a, b)
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mazamamaths.Add(1, 1)
	}
}
