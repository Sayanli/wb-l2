package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
Позволяет откладывать выполнение команд, выставлять их в очереди, хранить историю и делать отмену.
*/

import (
	"fmt"
)

type Command interface {
	Execute()
}

type ConcreteCommandA struct {
	receiver *Receiver
}

func (c *ConcreteCommandA) Execute() {
	c.receiver.Action("CommandA")
}

type ConcreteCommandB struct {
	receiver *Receiver
}

func (c *ConcreteCommandB) Execute() {
	c.receiver.Action("CommandB")
}

type Receiver struct{}

func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

type Invoker struct {
	history []Command
}

func (i *Invoker) StoreAndExecute(cmd Command) {
	i.history = append(i.history, cmd)
	for i, cmd := range i.history {
		fmt.Printf("history%d: ", i)
		cmd.Execute()
	}
}

/*
 func main() {
 	receiver := new(Receiver)
 	commandA := &ConcreteCommandA{receiver}
 	commandB := &ConcreteCommandB{receiver}
 	invoker := new(Invoker)
 	invoker.StoreAndExecute(commandA)
 	invoker.StoreAndExecute(commandB)
}
*/
