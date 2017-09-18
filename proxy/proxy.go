package main

import (
	"fmt"
)

type Image interface {
	Display()
}

type RealImage struct {
	fileName string
}

func (self *RealImage) Display() {
	fmt.Println("displaying " + self.fileName)
}

func (self *RealImage) loadFromDisk(fileName string) {
	fmt.Println("loading " + self.fileName + " from disk")
}

func NewRealImage(fileName string) *RealImage {
	realImage := &RealImage{fileName}
	realImage.loadFromDisk(fileName)
	return realImage
}

type ProxyImage struct {
	realImage *RealImage
	fileName  string
}

func (self *ProxyImage) Display() {
	if self.realImage == nil {
		self.realImage = NewRealImage(self.fileName)
	}
	self.realImage.Display()
}

func NewProxyImage(fileName string) *ProxyImage {
	proxyImage := &ProxyImage{fileName: fileName}
	proxyImage.fileName = fileName
	return proxyImage
}

func main() {
	image := NewProxyImage("yosemite.jpg")

	// image loaded from disk
	image.Display()

	// image not loaded from disk
	image.Display()
}
