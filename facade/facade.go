package main

import "fmt"

type Server struct{}

func (self *Server) Listen(port int) {
	fmt.Printf("listening to port %v\n", port)
}

type Logging struct{}

func (self *Logging) Aggregate() {
	fmt.Println("aggregate distributed logs")
}

type Monitoring struct{}

func (self *Monitoring) Subscribe(event string) {
	fmt.Printf("continously watch %s\n", event)
}

type CloudSystemFacade struct {
	webServer *Server
	log       *Logging
	monitor   *Monitoring
}

func NewCloudSystemFacade() *CloudSystemFacade {
	return &CloudSystemFacade{&Server{}, &Logging{}, &Monitoring{}}
}

func (self *CloudSystemFacade) start() {
	self.webServer.Listen(80)
	self.log.Aggregate()
	self.monitor.Subscribe("network traffic")
	self.monitor.Subscribe("cpu resource")
}

func main() {
	cloudSystem := NewCloudSystemFacade()
	cloudSystem.start()
}
