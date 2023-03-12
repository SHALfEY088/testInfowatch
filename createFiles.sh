#!/bin/bash

# Указываем папку для создания файлов и количество файлов, которые нужно создать
folder_path="folderWithFiles"
n=10

# Создаем папку, если ее нет
mkdir -p "$folder_path"

# Создаем n файлов со случайным содержимым
for i in $(seq 1 $n); do
  # Генерируем случайное имя файла
  file_name=$(cat /dev/urandom | tr -dc 'a-z' | fold -w 10 | head -n 1)
  # Генерируем случайное содержимое файла
  file_content=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 100 | head -n 1)
  # Создаем файл в указанной папке и записываем в него содержимое
  echo "$file_content" > "$folder_path/$file_name.txt"
done
