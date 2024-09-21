package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != "POST" {
			http.Error(w, "Только POST запросы разрешены", http.StatusMethodNotAllowed)
			return
		}

		// Получаем файл из запроса
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Ошибка чтения файла", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Создаем файл на сервере
		out, err := os.Create(header.Filename)
		if err != nil {
			http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		// Копируем содержимое загруженного файла в новый файл
		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Файл %s успешно загружен", header.Filename)
	})

	http.ListenAndServe("localhost:8080", nil)
}
