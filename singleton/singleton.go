package main

import "fmt"

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()

	fmt.Println(instance1 == instance2)
}
