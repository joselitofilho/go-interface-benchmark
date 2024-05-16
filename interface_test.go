package main

import "testing"

func BenchmarkInterfaceCallSimple(b *testing.B) {
	var z zero
	var g getter = z

	b.Run("ViaInterface", func(b *testing.B) {
		total := 0
		for i := 0; i < b.N; i++ {
			total += g.get()
		}
		if total > 0 {
			b.Logf("total is %d", total)
		}
	})

	b.Run("Direct", func(b *testing.B) {
		total := 0
		for i := 0; i < b.N; i++ {
			total += z.get()
		}
		if total > 0 {
			b.Logf("total is %d", total)
		}
	})
}
