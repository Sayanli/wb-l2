package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackStr(s string) (string, error) {
	var last rune
	var sb strings.Builder

	for _, r := range s {
		if unicode.IsDigit(r) {
			if unicode.IsDigit(last) || last == 0 {
				err := fmt.Sprintf("the string '%s' isn't correct", s)
				return s, errors.New(err)
			}
			if repeat := int(r - '1'); repeat >= 0 {
				str := strings.Repeat(string(last), repeat)
				sb.WriteString(str)
			} else {
				return s, errors.New("don`t use '0' in the string")
			}
		} else {
			sb.WriteRune(r)
		}
		last = r
	}

	return sb.String(), nil
}

func main() {
	s, err := UnpackStr(``)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
