package main

import "fmt"

type Memento struct {
	state string
}

type Originator struct {
	state string
}

func (self *Originator) SaveState() *Memento {
	return &Memento{self.state}
}

func (self *Originator) GetState(memento *Memento) {
	self.state = memento.state
}

type CareTaker struct {
	mementoList []*Memento
}

func (self *CareTaker) Add(state *Memento) {
	self.mementoList = append(self.mementoList, state)
}

func (self *CareTaker) Get(index int) *Memento {
	return self.mementoList[index]
}

func main() {
	originator := &Originator{}
	careTaker := &CareTaker{}

	originator.state = "State 1"
	originator.state = "State 2"
	careTaker.Add(originator.SaveState())

	originator.state = "State 3"
	careTaker.Add(originator.SaveState())

	originator.state = "State 4"
	fmt.Println("Current state: ", originator.state)

	originator.GetState(careTaker.Get(0))
	fmt.Println("First saved state: ", originator.state)

	originator.GetState(careTaker.Get(1))
	fmt.Println("Second saved state: ", originator.state)
}
