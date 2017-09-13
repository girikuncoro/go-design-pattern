package main

import "fmt"

type WhatsappLegacy struct{}

func (self *WhatsappLegacy) SendMessageOldWay() string {
	return "send message whatsapp old way"
}

type MessageAdapter interface {
	SendMessage() string
}

type Adapter struct {
	whatsappLegacy *WhatsappLegacy
}

func WhatsappAdapter() *Adapter {
	return &Adapter{&WhatsappLegacy{}}
}

func (self *Adapter) SendMessage() string {
	fmt.Println("calling from adapter")
	return self.whatsappLegacy.SendMessageOldWay()
}

func main() {
	whatsapp := WhatsappAdapter()
	fmt.Println(whatsapp.SendMessage())
}
