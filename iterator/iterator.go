package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Container interface {
	GetIterator() Iterator
}

type SchoolIterator struct {
	schoolRepo *SchoolRepo
	index      int
}

func (self *SchoolIterator) HasNext() bool {
	if self.index < len(self.schoolRepo.schools) {
		return true
	}
	return false
}

func (self *SchoolIterator) Next() interface{} {
	defer func() { self.index++ }()
	if self.HasNext() {
		return self.schoolRepo.schools[self.index]
	}
	return nil
}

type SchoolRepo struct {
	schools []string
}

func (self *SchoolRepo) GetIterator() Iterator {
	return &SchoolIterator{schoolRepo: self}
}

func main() {
	schoolsRepo := &SchoolRepo{[]string{"Cornell", "Stanford", "MIT", "Harvard"}}
	for i := schoolsRepo.GetIterator(); i.HasNext(); {
		school := i.Next()
		fmt.Println(school)
	}
}
