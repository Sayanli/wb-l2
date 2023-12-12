package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fFields    string
	fDelimiter string
	fSeparated bool
)

func initFlags() {
	flag.StringVar(&fFields, "f", "", "select only these fields; input example: '-f=1,3', '-f=1'; also print any  line  that  contains no delimiter character, unless the -s option is specified")
	flag.StringVar(&fDelimiter, "d", "	", "use DELIM instead of TAB for field delimiter")
	flag.BoolVar(&fSeparated, "s", false, "don't print lines not containing delimiters")

	flag.Parse()
}

func main() {
	initFlags()
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()

		if !strings.Contains(line, fDelimiter) {
			fmt.Println(line)
			continue
		}

		parts := strings.Split(line, fDelimiter)

		for _, num := range strings.Split(fFields, ",") {
			i, _ := strconv.Atoi(num)
			fmt.Println(parts[i-1])

			if i < len(parts) {
				fmt.Println(fDelimiter)
			}
		}
	}

}
