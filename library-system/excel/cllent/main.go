package main

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Создание нового Excel файла
	f := excelize.NewFile()

	// Добавление данных в файл
	index, _ := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", "A1", "Hello")
	f.SetCellValue("Sheet1", "B1", "World")
	f.SetActiveSheet(index)

	//// Сохранение файла на диск
	//filePath := "file.xlsx"
	//if err := f.SaveAs(filePath); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// Отправка файла на сервер
	//err := sendExcelFile(filePath)
	//if err != nil {
	//	fmt.Println("Ошибка отправки файла:", err)
	//}

	err := sendX(f)
	if err != nil {
		fmt.Println("Ошибка отправки файла:", err)
	}
}

func sendX(f *excelize.File) error {
	// Сохранение Excel файла в память (буфер)
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return err
	}

	// Создание буфера для формы
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Создание файла в multipart форме
	part, err := writer.CreateFormFile("file", "file.xlsx")
	if err != nil {
		return err
	}

	// Копирование содержимого файла в форму
	_, err = part.Write(buf.Bytes())
	if err != nil {
		return err
	}

	// Закрытие формы
	err = writer.Close()
	if err != nil {
		return err
	}

	// Создание HTTP POST запроса
	url := "http://localhost:8080" // Замените на ваш URL
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// Установка заголовков
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверка ответа от сервера
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось отправить файл: статус %v", resp.Status)
	}

	fmt.Println("Файл успешно отправлен")
	return nil
}

func sendExcelFile(filePath string) error {
	// Открытие файла
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Создание буфера для формы
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Создание файла в multipart форме
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}

	// Копирование содержимого файла в форму
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Закрытие формы
	err = writer.Close()
	if err != nil {
		return err
	}

	// Создание HTTP POST запроса
	url := "http://localhost:8080" // Замените на URL вашего сервера
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return err
	}

	// Установка заголовков
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверка ответа от сервера
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось отправить файл: статус %v", resp.Status)
	}

	fmt.Println("Файл успешно отправлен")
	return nil
}
