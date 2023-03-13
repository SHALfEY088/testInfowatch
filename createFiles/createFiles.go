package createFiles

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"sync"
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

	// Используем каналы для передачи данных между горутинами
	fileNameChan := make(chan string, n)
	fileContentChan := make(chan string, n)

	// Получаем количество доступных процессоров
	numCPUs := runtime.NumCPU()

	// Используем wait group для ожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(numCPUs)

	// Создаем горутины, каждая из которых создает файл с заданным именем и содержимым
	for i := 0; i < numCPUs; i++ {
		// Генерируем случайное имя файла и отправляем его в канал fileNameChan
		go func() {
			fileName := generateRandomString(10)
			fileNameChan <- fileName
		}()

		// Генерируем случайное содержимое файла и отправляем его в канал fileContentChan
		go func() {
			fileContent := generateRandomStringWithSpace(100)
			fileContentChan <- fileContent
		}()

		// Создаем файл с заданным именем и содержимым, используя данные из каналов
		go func() {
			fileName := <-fileNameChan
			fileContent := <-fileContentChan

			// Создаем файл в указанной папке и записываем в него содержимое
			err := ioutil.WriteFile(fmt.Sprintf("%s/%s.txt", folderPath, fileName), []byte(fileContent), os.ModePerm)
			if err != nil {
				panic(err)
			}

			// Уменьшаем счетчик wait group
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// Генерирует случайную строку указанной длины
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Генерирует случайную строку указанной длины c пробелами и символами перевода строки
func generateRandomStringWithSpace(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 \n"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
