package main

import "fmt"

func main() {
	months := [...]string{
		1:  "一月",
		2:  "二月",
		3:  "三月",
		4:  "四月",
		5:  "五月",
		6:  "六月",
		7:  "七月",
		8:  "八月",
		9:  "九月",
		10: "十月",
		11: "十一月",
		12: "十二月",
	}

	fmt.Println(months) // months[0] == ""
	fmt.Println(months[:])
	fmt.Println(months[1:13])
	fmt.Println(months[1:4])

	summer := months[6:9]
	fmt.Println(summer)

	// 延长slice
	anotherSummer := summer[:5]
	fmt.Println(anotherSummer)

	// slice无法用==比较

	// slice上的操作会影响到原数组
	summer[0] = "June"
	fmt.Println(summer, months)
}
