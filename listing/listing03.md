Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

```
dynamicTypeInfo - ссылка на реализацию структуры, которая реализует данный интерфейс (например *main.MyBeautifulIntf)

dynamicValue - значение произвольного типа. unsafe.Pointer казывает на ячейку, где лежит значение, и читает данные из дескриптора типа для определения типа.

Собственно, если интерфейс не является пустым, то он никогда не будет равен nil, т.к. хранит в себе dynamicTypeInfo с ссылкой на реализацию. Только пустой интерфейс - interface{} будет равен nil.