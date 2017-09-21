package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	data string
}

func (self *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, self.data) {
		return true
	}
	return false
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (self *OrExpression) Interpret(context string) bool {
	return self.expr1.Interpret(context) || self.expr2.Interpret(context)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (self *AndExpression) Interpret(context string) bool {
	return self.expr1.Interpret(context) && self.expr2.Interpret(context)
}

func main() {
	// Example: Donald and Trump are male
	donald := &TerminalExpression{"Donald"}
	trump := &TerminalExpression{"Trump"}
	isMale := &OrExpression{donald, trump}

	// Example: Natalie is a single lady
	natalie := &TerminalExpression{"Natalie"}
	single := &TerminalExpression{"Single"}
	isSingleLady := &AndExpression{natalie, single}

	fmt.Println("Donald is male?", isMale.Interpret("Donald"))
	fmt.Println("Natalie is single lady?", isSingleLady.Interpret("Single Natalie"))

}
