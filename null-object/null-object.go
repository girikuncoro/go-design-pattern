package main

import "fmt"

type User interface {
	IsNil() bool
	GetName() string
}

type RealUser struct {
	name string
}

func (self *RealUser) IsNil() bool {
	return false
}

func (self *RealUser) GetName() string {
	return self.name
}

type NilUser struct{}

func (self *NilUser) IsNil() bool {
	return true
}

func (self *NilUser) GetName() string {
	return "User not found in database"
}

type UserFactory struct {
	nameList []string
}

func (self *UserFactory) GetUser(username string) User {
	for _, name := range self.nameList {
		if name == username {
			return &RealUser{name: username}
		}
	}
	return &NilUser{}
}

func main() {
	userFactory := &UserFactory{[]string{"Allie", "Bob", "David"}}
	users := []string{"Allie", "Bob", "Charlie", "David", "Edward"}

	for _, user := range users {
		fmt.Println(userFactory.GetUser(user).GetName())
	}
}
