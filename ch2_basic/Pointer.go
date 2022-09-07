package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	var a int = 0
	incr(&a)
	fmt.Println(incr(&a))
}
