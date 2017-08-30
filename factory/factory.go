package main

import "fmt"

type Coffee interface {
	Brew() string
}

type EspressoCoffee struct{}

func (espresso *EspressoCoffee) Brew() string {
	return "Pour espresso coffee"
}

type MacchiatoCoffee struct{}

func (macchiato *MacchiatoCoffee) Brew() string {
	return "Pour macchiato coffee"
}

type CapuccinoCoffee struct{}

func (capuccino *CapuccinoCoffee) Brew() string {
	return "Pour capuccino coffee"
}

type CoffeeType int

const (
	Espresso CoffeeType = 1 << iota
	Macchiato
	Capuccino
)

func CoffeeFactory(t CoffeeType) Coffee {
	switch t {
	case Espresso:
		return &EspressoCoffee{}
	case Macchiato:
		return &MacchiatoCoffee{}
	case Capuccino:
		return &CapuccinoCoffee{}
	default:
		return nil
	}
}

func main() {
	coffees := []Coffee{
		CoffeeFactory(Espresso),
		CoffeeFactory(Macchiato),
		CoffeeFactory(Capuccino),
	}

	for _, coffee := range coffees {
		fmt.Println(coffee.Brew())
	}
}
