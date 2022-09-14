package main

import (
	"fmt"
	"goLearn/ch7_interface/bytecounter"
	"sort"
)

func main() {
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "haunyu"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	// 任何值都可以赋给空接口
	var any interface{}
	any = true
	any = 12.334
	// 实际上 Println 接收的也是空接口参数，这样它才能打印任何类型的数据
	// 行为上和泛型比较接近
	fmt.Println(any)

	var str = [3]string{"123", "342", "001"}
	ss := str[1:]
	sort.Sort(sort.StringSlice(ss))
	fmt.Println(ss)

}
