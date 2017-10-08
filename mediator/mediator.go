package main

type Mediator interface {
	SendMessage()
	ShowMessage()
}

type ChatMediator struct{}

type User struct{}

type Chatroom struct{}

func main() {

}
