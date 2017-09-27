package main

import "fmt"

type Strategy interface {
	Solve(num1 int, num2 int) int
}

type Add struct{}

func (self *Add) Solve(num1 int, num2 int) int {
	return num1 + num2
}

type Subtract struct{}

func (self *Subtract) Solve(num1 int, num2 int) int {
	return num1 - num2
}

type Multiply struct{}

func (self *Multiply) Solve(num1 int, num2 int) int {
	return num1 * num2
}

type Context struct {
	strategy Strategy
}

func (self *Context) ExecuteStrategy(num1 int, num2 int) int {
	return self.strategy.Solve(num1, num2)
}

func main() {
	context := &Context{&Add{}}
	fmt.Println(context.ExecuteStrategy(100, 4))

	context = &Context{&Subtract{}}
	fmt.Println(context.ExecuteStrategy(100, 4))

	context = &Context{&Multiply{}}
	fmt.Println(context.ExecuteStrategy(100, 4))
}
