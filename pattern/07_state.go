package main

import "fmt"

type State interface {
	doAction(context Context)
}

type StartState struct{}

type StopState struct{}

type Context struct {
	state State
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) getState() State {
	return c.state
}

func (c *Context) doAction() {
	c.state.doAction(*c)
}

func (s *StartState) doAction(context Context) {
	fmt.Println("Player is in start state")
	context.setState(s)
}

func (s *StopState) doAction(context Context) {
	fmt.Println("Player is in stop state")
	context.setState(s)
}

func main() {
	context := Context{}
	startState := StartState{}
	stopState := StopState{}
	context.setState(&startState)
	context.doAction()
	context.setState(&stopState)
	context.doAction()
}
