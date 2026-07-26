package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/RGBFox/Final-6-sprint/internal/service"
)

func MainHandle(w http.ResponseWriter, req *http.Request) {

	// открываем файл
	root, err := os.OpenRoot(".")
	if err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()
	// Открываем файл для отправки

	file, err := root.Open("index.html")
	if err != nil {
		http.Error(w, "файл /index.html не найден", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// копируем содержимое файла в ответ
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "ошибка при отправке файла", http.StatusInternalServerError)
		return
	}
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "ошибка при парсинге", http.StatusInternalServerError)
		return
	}

	// получаем файл из формы
	file, head, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Читаем файл и конвертируем
	needtext, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Невозможно прочитать файл", http.StatusInternalServerError)
		return
	}
	convert, err := service.Convert(string(needtext))
	if err != nil {
		http.Error(w, "Невозможно конвертировать файл", http.StatusInternalServerError)
		return
	}

	name := time.Now().UTC().Format("2006-01-02_15-04-05")

	// Создаем файл с конвертацией
	newFile, err := os.OpenFile(name+filepath.Ext(head.Filename), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	_, err = newFile.WriteString(convert)
	if err != nil {
		http.Error(w, "записи файла", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, convert)
}
