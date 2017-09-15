package main

import "fmt"

type AbstractFile interface {
	GetName() string
	PrintList()
}

type File struct {
	name string
}

func (self *File) GetName() string {
	return self.name
}

func (self *File) PrintList() {
	fmt.Println(self.name)
}

type Directory struct {
	includedFiles []AbstractFile
	name          string
}

func (self *Directory) GetName() string {
	return self.name
}

func (self *Directory) Add(entry AbstractFile) {
	self.includedFiles = append(self.includedFiles, entry)
}

func (self *Directory) PrintList() {
	fmt.Println()
	for _, file := range self.includedFiles {
		fmt.Print(self.GetName() + "/")
		file.PrintList()
	}
}

func main() {
	music := &Directory{name: "MUSIC"}
	johnMayer := &Directory{name: "JOHN-MAYER"}
	incubus := &Directory{name: "INCUBUS"}
	track1 := &File{name: "free falling.mp3"}
	track2 := &File{name: "anna molly.mp3"}
	track3 := &File{name: "californication.mp3"}
	track4 := &File{name: "drive.mp3"}
	track5 := &File{name: "dig.mp3"}
	track6 := &File{name: "xoxo.mp3"}
	music.Add(track3)
	music.Add(johnMayer)
	music.Add(incubus)
	johnMayer.Add(track1)
	johnMayer.Add(track4)
	johnMayer.Add(track6)
	incubus.Add(track2)
	incubus.Add(track5)

	music.PrintList()
}
