package countSymbols

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// CountSymbolsInFiles подсчет символов в файлах
func CountSymbolsInFiles() {
	var mutex sync.Mutex
	// Получаем количество доступных процессоров
	numCPUs := runtime.NumCPU()

	// Создаем WaitGroup для синхронизации завершения всех горутин
	wg := &sync.WaitGroup{}
	wg.Add(numCPUs)

	// Создаем мапу для хранения количества встречающихся символов
	charCount := make(map[rune]int)

	// Указываем путь к папке с файлами
	folderPath := "folderWithFiles"

	// Открываем папку и перебираем все файлы в ней
	files, err := os.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Println("В папке .folderWithFiles/ нет файлов")
	} else {
		fmt.Printf("В указанной папке есть %d файлов, вот гистограмма распределения символов в них:\n", len(files))
	}

	// Создаем канал для передачи файлов в горутины
	fileChan := make(chan os.DirEntry)

	// Запускаем несколько горутин для обработки файлов
	for i := 0; i < numCPUs; i++ {
		go func() {
			defer wg.Done()

			// Обрабатываем файлы из канала
			for file := range fileChan {
				// Открываем каждый файл и считываем его содержимое
				fileContent, err := os.Open(filepath.Join(folderPath, file.Name()))
				if err != nil {
					panic(err)
				}

				scanner := bufio.NewScanner(fileContent)
				buf := bytes.Buffer{}

				for scanner.Scan() {
					buf.WriteString(scanner.Text())
				}

				// Считаем количество встречающихся символов и добавляем в общую мапу
				for _, r := range buf.String() {
					mutex.Lock()
					charCount[r]++
					mutex.Unlock()
				}

				fileContent.Close()
			}
		}()
	}

	// Передаем файлы в канал для обработки горутинами
	for _, file := range files {
		fileChan <- file
	}
	close(fileChan)

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим гистограмму распределения символов
	for char, count := range charCount {
		fmt.Printf("%c %d\n", char, count)
	}
}
