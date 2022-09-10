package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var a Employee
	fmt.Println(a)

	var b *Employee = &a
	b.ID = 100
	fmt.Println(a, *b)

	aTree := tree{100, nil, nil}
	bTree := tree{1, &aTree, nil}
	fmt.Println(aTree, bTree)
	fmt.Println(bTree.left.value)

	// 结构体可比较
	fmt.Println(aTree == bTree, aTree == *bTree.left)
}
