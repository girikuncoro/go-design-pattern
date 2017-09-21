package main

import (
	"fmt"
	"strconv"
)

type Order interface {
	Execute()
}

type Stock struct {
	name   string
	amount int
}

func (self *Stock) Buy() {
	fmt.Println("buy " + self.name + " stock " + strconv.Itoa(self.amount) + " shares")
}

func (self *Stock) Sell() {
	fmt.Println("sell " + self.name + " stock " + strconv.Itoa(self.amount) + " shares")
}

type BuyStock struct {
	stock *Stock
}

func (self *BuyStock) Execute() {
	self.stock.Buy()
}

type SellStock struct {
	stock *Stock
}

func (self *SellStock) Execute() {
	self.stock.Sell()
}

// command invoker which is going to
// take and execute request
type Broker struct {
	orderList []Order
}

func (self *Broker) TakeOrder(order Order) {
	self.orderList = append(self.orderList, order)
}

func (self *Broker) PlaceOrders() {
	for _, order := range self.orderList {
		order.Execute()
	}
	self.orderList = nil
}

func main() {
	stock := &Stock{name: "APPL", amount: 10}
	buyStock := &BuyStock{stock}
	sellStock := &SellStock{stock}

	broker := &Broker{}
	broker.TakeOrder(buyStock)
	broker.TakeOrder(sellStock)

	broker.PlaceOrders()
}
