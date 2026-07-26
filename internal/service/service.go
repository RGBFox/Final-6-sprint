// service - содержит Функцию автоматического определения кода Морзе или обычного текста из переданной строки. Если передан обычный текст, функция должна переконвертировать его в код Морзе и вернуть; и наоборот — если был передан код Морзе, функция должна переконвертировать его в обычный текст и вернуть.
package service

import (
	"fmt"
	"strings"

	"github.com/RGBFox/Final-6-sprint/pkg/morse"
)

// Alpha содержит все букевы и числа русского языка
const Alpha string = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ0123456789"

// Convert - конвертирует все данные в морзянку или наоборот в зависимости от полученныз данных
func Convert(conv string) (string, error) {
	if len(conv) < 1 {
		return "", fmt.Errorf("нельзя отправлять пустую строку")
	}
	if strings.ContainsAny(conv, Alpha) {
		return morse.ToMorse(conv), nil
	}
	return morse.ToText(conv), nil

}
