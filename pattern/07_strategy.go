package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Плюсы
Принцип открытости/закрытости.
Изолирует код и данные алгоритмов от остальных классов.
Уход от наследования к делегированию.
Минусы
Усложняет программу за счёт дополнительных классов.
Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type Context struct {
	strategy func()
}

func (c *Context) Execute() {
	c.strategy()
}

func (c *Context) SetStrategy(strategy func()) {
	c.strategy = strategy
}

/*
func main() {
 	concreteStrategyA := func() {
 		fmt.Println("concreteStrategyA()")
 	}
 	concreteStrategyB := func() {
 		fmt.Println("concreteStrategyB()")
 	}
 	context := Context{concreteStrategyA}
 	context.Execute()
 	context.SetStrategy(concreteStrategyB)
 	context.Execute()
}
*/
