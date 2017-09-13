package main

import "fmt"

type DrawAPI interface {
	DrawCircle(int, int, int)
}

type RedCircle struct{}

func (self *RedCircle) DrawCircle(radius int, x int, y int) {
	fmt.Printf("drawing red circle, radius: %v, x: %v, y: %v\n", radius, x, y)
}

type GreenCircle struct{}

func (self *GreenCircle) DrawCircle(radius int, x int, y int) {
	fmt.Printf("drawing green circle, radius: %v, x: %v, y: %v\n", radius, x, y)
}

type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius int
	draw         DrawAPI
}

func (self *Circle) Draw() {
	self.draw.DrawCircle(self.radius, self.x, self.y)
}

func main() {
	redCircle := &Circle{100, 100, 10, &RedCircle{}}
	greenCircle := &Circle{100, 100, 10, &GreenCircle{}}

	redCircle.Draw()
	greenCircle.Draw()
}
