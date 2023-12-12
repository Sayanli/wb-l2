package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Плюсы
Избавляет класс от привязки к конкретным классам продуктов.
Выделяет код производства продуктов в одно место, упрощая поддержку кода.
Упрощает добавление новых продуктов в программу.
Реализует принцип открытости/закрытости.
Минусы
Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
*/

import (
	"fmt"
)

type Creator struct {
	factory factory
}

func (c *Creator) Operation() {
	product := c.factory.factoryMethod()
	product.method()
}

type factory interface {
	factoryMethod() Product
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) factoryMethod() Product {
	return new(ConcreteProduct)
}

type Product interface {
	method()
}

type ConcreteProduct struct{}

func (p *ConcreteProduct) method() {
	fmt.Println("ConcreteProduct.method()")
}

/*
 func main() {
 	creator := Creator{new(ConcreteCreator)}
 	creator.Operation()
}
*/
