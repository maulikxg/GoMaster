package test

import "testing"

func FuzzEqual(f *testing.F) {

	f.Add([]byte{'a', 'b', 'c', 'd'}, []byte{'e', 'f', 'g', 'h'})

	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})

}
