package main

import "testing"

type benchset struct {
	name string
	f    benchFunc
}
type benchFunc = func() (string, []any)

func bench(b *testing.B, f benchFunc) {
	for i := 0; i < b.N; i++ {
		f()
	}
}
