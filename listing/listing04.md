Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
...
9
fatal error: all goroutines are asleep - deadlock!
```

Дэдлок из за того, что осталась одна заблокированная горутина - main-функция, которая читает из канала. Достаточно добавить закрытие канала в горутине после цикла с записью в канал.
