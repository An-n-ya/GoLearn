package main

import "fmt"

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func cToF(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	const freezingC, boilingC = 0.0, 100.0
	fmt.Printf("%g℃ = %g℉ \n", freezingC, cToF(freezingC))
	fmt.Printf("%g℃ = %g℉ \n", boilingC, cToF(boilingC))
}
