package main

import "fmt"

type TaxiElement interface {
	Accept(VisitorPassenger)
}

type Uber struct {
	name string
}

func (self *Uber) Accept(visitor VisitorPassenger) {
	self.name = "Uber"
	visitor.CallUber(self)
}

type Lyft struct {
	name string
}

func (self *Lyft) Accept(visitor VisitorPassenger) {
	self.name = "Lyft"
	visitor.CallLyft(self)
}

type Gojek struct {
	name string
}

func (self *Gojek) Accept(visitor VisitorPassenger) {
	self.name = "Gojek"
	visitor.CallGojek(self)
}

type VisitorPassenger interface {
	CallUber(*Uber)
	CallLyft(*Lyft)
	CallGojek(*Gojek)
}

type BusinessmanPassenger struct{}

func (self *BusinessmanPassenger) CallUber(uber *Uber) {
	fmt.Println("Businessman is riding " + uber.name)
}

func (self *BusinessmanPassenger) CallLyft(lyft *Lyft) {
	fmt.Println("Businessman is riding " + lyft.name)
}

func (self *BusinessmanPassenger) CallGojek(gojek *Gojek) {
	fmt.Println("Businessman is riding " + gojek.name)
}

type StudentPassenger struct{}

func (self *StudentPassenger) CallUber(uber *Uber) {
	fmt.Println("Student is riding " + uber.name)
}

func (self *StudentPassenger) CallLyft(lyft *Lyft) {
	fmt.Println("Student is riding " + lyft.name)
}

func (self *StudentPassenger) CallGojek(gojek *Gojek) {
	fmt.Println("Student is riding " + gojek.name)
}

func main() {
	taxis := []TaxiElement{&Uber{}, &Lyft{}, &Gojek{}}
	businessman := &BusinessmanPassenger{}
	student := &StudentPassenger{}

	for _, taxi := range taxis {
		taxi.Accept(businessman)
	}
	for _, taxi := range taxis {
		taxi.Accept(student)
	}
}
