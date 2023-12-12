package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("P$ ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		switch cmdName := args[0]; cmdName {
		case "exit":
			return
		case "cd":
			if len(args) < 2 {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println(err)
				}
				args = append(args, homeDir)
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "fork":
			cmd := exec.Command(os.Args[0], append(os.Args[1:], "forked")...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			cmd := exec.Command(cmdName, args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
