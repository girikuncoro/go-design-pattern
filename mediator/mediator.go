package main

import "fmt"

type Mediator interface {
	AddUser(User)
	SendMessage(User, string)
}

type Chatroom struct {
	users []*User
}

func (self *Chatroom) AddUser(user *User) {
	self.users = append(self.users, user)
}

func (self *Chatroom) SendMessage(sender *User, message string) {
	for _, user := range self.users {
		if user != sender {
			fmt.Println(sender.name + " sent message to " + user.name + ": " + message)
		}
	}
}

type User struct {
	name     string
	mediator *Chatroom
}

func (self *User) Send(message string) {
	self.mediator.SendMessage(self, message)
}

func NewUser(name string, mediator *Chatroom) *User {
	user := &User{name, mediator}
	mediator.AddUser(user)
	return user
}

func main() {
	chatroom := &Chatroom{}
	larry := NewUser("Larry", chatroom)
	mark := NewUser("Mark", chatroom)
	bill := NewUser("Bill", chatroom)

	larry.Send("Hello from Larry!")
	mark.Send("Aloha from Mark!!")
	bill.Send("Hi from Bill!!!")
}
