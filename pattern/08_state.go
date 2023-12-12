package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Плюсы
Избавляет от множества больших условных операторов машины состояний.
Концентрирует в одном месте код, связанный с определённым состоянием.
Упрощает код контекста.
Минусы
Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

import (
	"fmt"
)

type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle()
}

func (c *Context) SetState(state State) {
	c.state = state
}

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}

/*
func main() {
 	context := Context{new(ConcreteStateA)}
 	context.Request()
 	context.SetState(new(ConcreteStateB))
 	context.Request()
}
*/
