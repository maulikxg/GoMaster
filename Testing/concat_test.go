package test

import (
	"testing"
	"time"
)

func BenchmarkJoinStrings(b *testing.B) {

	strs := []string{"Hello", ",", "world", "!"}

	// The benchmark runner will call this function b.N times
	//for i := 0; i < b.N; i++ {
	//
	//}

	for i := 0; i < 700; i++ {
		strs = append(strs, "ababba")
	}

	time.Sleep(time.Second * 3)

	b.ReportAllocs()

	for b.Loop() {
		JoinStrings(strs)
	}

}
