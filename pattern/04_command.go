package main

import "fmt"

type Command interface{
	Execute()
}

type FirstCommand struct{
	receiver * Receiver
}

func (c *FirstCommand) Execute(){
	c.receiver.Action()
}

type SecondCommand struct{
	receiver * Receiver
}

func (c *SecondCommand) Execute(){
	c.receiver.Action2()
}

type Invoker struct{
	command Command
}

func (i *Invoker) SetCommand(command Command){
	i.command = command
}

func (i *Invoker) ExecuteCommand(){
	i.command.Execute()
}

type Receiver struct{}

func (r *Receiver) Action(){
	fmt.Println("Action")
}

func (r *Receiver) Action2(){
	fmt.Println("Action2")
}



func main(){
	receiver := new(Receiver)
	command1 := new(FirstCommand)
	command1.receiver = receiver
	command2 := new(SecondCommand)
	command2.receiver = receiver

	invoker := new(Invoker)
	invoker.SetCommand(command1)
	invoker.ExecuteCommand()
	invoker.SetCommand(command2)
	invoker.ExecuteCommand()
}