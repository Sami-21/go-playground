package main

func adder() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}
