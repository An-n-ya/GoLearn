package main

import (
	"fmt"
	"goLearn/ch3_datatype/printints"
	"strconv"
)

func main() {
	x, err := strconv.Atoi("123")
	if err != nil {
		return
	}
	fmt.Println(x)
	fmt.Println(printints.IntsToString([]int{1, 2, 3}))
}
