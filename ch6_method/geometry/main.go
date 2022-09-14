package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Distance 这里的p就相当于java中的this， Python中的self （接收名）
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(s float64) {
	p.X *= s
	p.Y *= s
}

func main() {
	p := Point{10, 2}
	q := Point{2, 8}
	fmt.Println(p.Distance(q))
	p.ScaleBy(10)
	(&p).ScaleBy(0.1)
	fmt.Println(p)

	pptr := &p
	fmt.Println(pptr.Distance(q))

	// coloredPoint 可以使用 Point 下的所有方法
	c_p := ColoredPoint{Point{2, 3}, color.RGBA{G: 255, A: 255}}
	c_q := ColoredPoint{Point{5, 8}, color.RGBA{R: 255, A: 255}}
	c_q.ScaleBy(3)
	// 结构体嵌套没有多态, 也就是说ColoredPoint不能直接当做Point使用
	fmt.Println(c_q, c_p)
	fmt.Println(c_p.Distance(c_q.Point))
	// c_p.Distance(c_q) 这是错误的

	// 匿名结构
	var line = struct {
		Start Point
		End   Point
	}{
		Point{1, 2},
		Point{2, 4},
	}
	fmt.Println(line)

	// 方法可以赋值给变量，重复利用
	disFromP := p.Distance
	fmt.Println(disFromP(q))

	// 方法可以抽象成函数
	// 这种语法称作 方法表达式
	distance := Point.Distance
	fmt.Println(distance(p, q))

}
