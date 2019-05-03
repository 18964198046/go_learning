package fib

import (
	"testing"
)

var a int
var b int

func TestFibList(t *testing.T) {
	a = 1
	b = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
