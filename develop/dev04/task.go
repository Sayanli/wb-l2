package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func anagram(words *[]string) *map[string]*[]string {
	anagrams := make(map[string][]string)

	for _, word := range *words {
		if len(word) <= 1 {
			continue
		}
		word := strings.ToLower(word)
		letters := strings.Split(word, "")
		sort.Strings(letters)
		key := strings.Join(letters, "")

		if _, ok := anagrams[key]; !ok {
			// 3
			anagrams[key] = []string{word}
		} else {
			anagrams[key] = append(anagrams[key], word)
		}
	}

	result := make(map[string]*[]string)

	for _, words := range anagrams {
		key := words[0]
		sl := make([]string, 0, len(words))
		result[key] = &sl

		sort.Strings(words)
		var last string
		for _, word := range words {
			if word == last {
				continue
			}
			last = word
			*result[key] = append(*result[key], word)
		}
	}

	return &result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "одуван"}

	am := *anagram(&words)
	fmt.Println(am)

	for key, value := range am {
		fmt.Println(key, *value)
	}
}
