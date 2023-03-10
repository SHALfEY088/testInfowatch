package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Создаем мапу для хранения количества встречающихся символов
	charCount := make(map[rune]int)

	// Указываем путь к папке с файлами
	folderPath := "folderWithFiles"

	// Открываем папку и перебираем все файлы в ней
	files, err := os.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		// Открываем каждый файл и считываем его содержимое
		fileContent, err := os.Open(filepath.Join(folderPath, file.Name()))
		if err != nil {
			panic(err)
		}
		defer fileContent.Close()

		scanner := bufio.NewScanner(fileContent)
		buf := bytes.Buffer{}

		for scanner.Scan() {
			buf.WriteString(scanner.Text())
		}

		for _, r := range buf.String() {
			charCount[r]++
		}
	}

	// Выводим гистограмму распределения символов
	for char, count := range charCount {
		fmt.Printf("%c %d\n", char, count)
	}
}
