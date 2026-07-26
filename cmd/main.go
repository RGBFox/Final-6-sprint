package main

import (
	"log"
	"os"

	"github.com/RGBFox/Final-6-sprint/internal/server"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal("ошибка создания файла")
	}
	defer file.Close()

	logger := log.New(file, "APP: ", log.LstdFlags)
	logger.Println("программа запущена")

	serv := server.Route(logger)
	err = serv.Htserver.ListenAndServe()
	if err != nil {
		logger.Fatal("ошибка создания сервера")
	}
}
