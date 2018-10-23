package main

import ()

const length = 1000000

func main() {
	a := make([]int, length, length)
	b := make([]int, length, length)
	c := make([]int, length, length)
	for i := 0; i < length; i++ {
		a[i] = i
		b[i] = i / 2
	}
	sum(a, b, c)
}

func sum(a, b, c []int) {
	for i := 0; i < length; i++ {
		c[i] = a[i] + b[i]
	}
}
