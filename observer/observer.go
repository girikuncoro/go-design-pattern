package main

import (
	"fmt"
	"strconv"
)

type Notifier struct {
	observers []Observer
	state     int
}

func (self *Notifier) SetState(state int) {
	fmt.Println("set state to: " + strconv.Itoa(state))
	self.state = state
	self.notify()
}

func (self *Notifier) Register(observer Observer) {
	self.observers = append(self.observers, observer)
}

func (self *Notifier) notify() {
	for _, observer := range self.observers {
		observer.update()
	}
}

type Observer interface {
	update()
}

type BinaryObserver struct {
	notifier *Notifier
}

func (self *BinaryObserver) update() {
	fmt.Println("binary string: " + strconv.FormatInt(int64(self.notifier.state), 2))
}

func NewBinaryObserver(notifier *Notifier) {
	binaryObserver := &BinaryObserver{}
	binaryObserver.notifier = notifier
	binaryObserver.notifier.Register(binaryObserver)
}

type HexaObserver struct {
	notifier *Notifier
}

func (self *HexaObserver) update() {
	fmt.Printf("hexadecimal string: %x\n", self.notifier.state)
}

func NewHexaObserver(notifier *Notifier) {
	hexaObserver := &HexaObserver{}
	hexaObserver.notifier = notifier
	hexaObserver.notifier.Register(hexaObserver)
}

func main() {
	notifier := &Notifier{}
	NewBinaryObserver(notifier)
	NewHexaObserver(notifier)

	notifier.SetState(9)
	notifier.SetState(4)
	notifier.SetState(123)
}
