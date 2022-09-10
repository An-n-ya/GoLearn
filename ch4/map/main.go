package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 12
	ages["charlie"] = 16
	fmt.Println(ages)

	ages2 := map[string]int{
		"alice":   12,
		"charlie": 16,
	}

	fmt.Println(ages2)

	// 不能比较
	//fmt.Println("ages == ages2: ", ages2 == ages)
	// 可以用自己实现的函数判断
	fmt.Println("ages == ages2", equal(ages, ages2))

	for name, age := range ages {
		fmt.Println(name, ":", age)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
