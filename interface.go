package main

type getter interface {
	get() int
}
type zero struct{}

//go:noinline
func (z zero) get() int {
	return 0
}
