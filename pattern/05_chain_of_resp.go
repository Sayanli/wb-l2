package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Плюсы
Маленькая зависимость между клиентом и обработчиками;
Реализует принцип единсвенной ответственности;
Реализует принцип открытости/закрытости;
Минусы
Запрос может остаться никам не обработаным - в таком случае такие запросы нужно логировать и потом разбирать, почему он не был обработан и нужно ли ввести новые обработчики для таких запросов, или выявить ошибку, если запрос некорректен.
*/

import (
	"fmt"
)

type Handler interface {
	Request(flag bool)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) Request(flag bool) {
	fmt.Println("ConcreteHandlerA.Request()")
	if flag {
		h.next.Request(flag)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) Request(flag bool) {
	fmt.Println("ConcreteHandlerB.Request()")
}

/*
func main() {
 	handlerA := &ConcreteHandlerA{new(ConcreteHandlerB)}
 	handlerA.Request(true)
}
*/
