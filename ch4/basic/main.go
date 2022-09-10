package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	q := [...]int{1, 2, 3}
	fmt.Printf("%T, %d\n", q, q[1])

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	fmt.Println("RMB:", RMB)

	r := [...]int{10: 1}
	fmt.Println("r:", r)

	aa := [2]int{1, 2}
	bb := [2]int{1, 2}
	fmt.Println("aa == bb: ", aa == bb)
}

// 清空固定长度的数组
// 输入只能是32长的byte数组
func clear(ptr *[32]byte) {
	*ptr = [32]byte{}
}
