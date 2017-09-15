package main

import (
	"fmt"
	"strconv"
)

type Gazillion struct {
	row int
}

func NewGazillion(row int) *Gazillion {
	fmt.Println("col to row: " + strconv.Itoa(row))
	return &Gazillion{row}
}

func (self *Gazillion) print(col int) {
	fmt.Print(" " + strconv.Itoa(self.row) + strconv.Itoa(col))
}

type Factory struct {
	pool map[int]*Gazillion
}

func NewFactory() *Factory {
	return &Factory{make(map[int]*Gazillion)}
}

func (self *Factory) GetGazillion(row int) *Gazillion {
	if self.pool[row] == nil {
		self.pool[row] = NewGazillion(row)
	}
	return self.pool[row]
}

func main() {
	ROWS := 6
	COLS := 10

	factory := NewFactory()
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			factory.GetGazillion(i).print(j)
		}
		fmt.Println()
	}
}
