package createFiles

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func CreateFiles() {
	// Указываем папку для создания файлов и количество файлов, которые нужно создать
	folderPath := "folderWithFiles"
	n := 10

	// Создаем папку, если ее нет
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем n файлов со случайным содержимым
	for i := 0; i < n; i++ {
		// Генерируем случайное имя файла
		fileName := generateRandomString(10)
		// Генерируем случайное содержимое файла
		fileContent := generateRandomString(100)
		// Создаем файл в указанной папке и записываем в него содержимое
		err := ioutil.WriteFile(fmt.Sprintf("%s/%s.txt", folderPath, fileName), []byte(fileContent), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

// Генерирует случайную строку указанной длины
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 \n"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
