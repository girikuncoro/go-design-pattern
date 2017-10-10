package main

import (
	"fmt"
	"math"
	"strconv"
)

type CircleData struct {
	x, y, radius int
	color        string
}

type Circle struct {
	circleData *CircleData
}

func NewCircle(x int, y int, radius int, color string) *Circle {
	circle := &Circle{
		&CircleData{x, y, radius, color},
	}
	return circle
}

func (self *Circle) Circumference() float64 {
	return float64(self.circleData.radius) * math.Pi
}

func (self *Circle) Diameter() int {
	return self.circleData.radius * 2
}

func (self *Circle) Draw() {
	fmt.Println("Draw circle with x: " + strconv.Itoa(self.circleData.x) + ", y: " + strconv.Itoa(self.circleData.y) + ", color: " + self.circleData.color)
}

func main() {
	circle := NewCircle(1, 1, 3, "red")
	fmt.Println(circle.Circumference())
	fmt.Println(circle.Diameter())
	circle.Draw()
}
