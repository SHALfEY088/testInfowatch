package main

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/SHALfEY088/testInfowatch/countSymbols"
	. "github.com/SHALfEY088/testInfowatch/createFiles"
)

func main() {
	var input string

	for {
		fmt.Println("\nВведите 1, чтобы создать файлы с случайным содержимым \nВведите 2, чтобы подсчитать количество символов в файлах\nВведите 3, чтобы выйти")

		fmt.Scanln(&input)

		switch input {
		case "1":
			CreateFiles()
			fmt.Println("Файлы успешно созданы!")
		case "2":
			CountSymbolsInFiles()
		case "3":
			fmt.Println("Программа завершена!")
			cmd := exec.Command("cmd", "/C", "pause")
			cmd.Stdout = os.Stdout
			cmd.Run()
			return
		default:
			fmt.Println("Неверный ввод, попробуйте снова")
		}
	}
}
