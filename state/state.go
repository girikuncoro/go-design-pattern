package main

import "fmt"

type State interface {
	Action(*Context)
	Move()
}

type Context struct {
	state State
}

type RunningState struct{}

func (self *RunningState) Action(context *Context) {
	fmt.Println("Robot is in running state")
	context.state = self
}

func (self *RunningState) Move() {
	fmt.Println("Robot runs 10 m/s")
}

type WalkingState struct{}

func (self *WalkingState) Action(context *Context) {
	fmt.Println("Robot is in walking state")
	context.state = self
}

func (self *WalkingState) Move() {
	fmt.Println("Robot walks 2 m/s")
}

func main() {
	context := &Context{}

	runningState := &RunningState{}
	runningState.Action(context)

	context.state.Move()

	walkingState := &WalkingState{}
	walkingState.Action(context)

	context.state.Move()
}
