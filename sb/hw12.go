package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//fmt.Println("========================")
	//fmt.Println("Задача 12.01.")
	//fmt.Println("========================")
	task1202()
	//fmt.Println("========================")
	//fmt.Println("Задача 12.02.")
	//fmt.Println("========================")
	//task1102()
}

func task1202()  {
	//Урок №2 Работа с файлами
	//Написать программу, которая на вход получала бы строку, введенную пользователем, а в файл писала: № строки,
	//дата и сообщение в формате:
	//2020-02-10 15:00:00 продам гараж
	//при вводе слова “выход” производился бы выход из программы
	fileName := "hw12.2.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Ошибка создания файла!", err)
		return
	}
	defer file.Close()

	var str, fullStr string
	strExit := "ВЫХОД"
	fmt.Println("Введите сообщения или команду 'Выход':")
	timeFormat := "2006-01-02 15:04:05"
	scanStdin := bufio.NewScanner(os.Stdin)
	for {
		scanStdin.Scan()
		str = scanStdin.Text()
		str = strings.Trim(str, " ")
		if strings.ToUpper(str) == strExit {
			break
		}
		fullStr += time.Now().Format(timeFormat) + " " + str + "\n"
		//fmt.Println(str)
	}
	fmt.Println("Работа программы завершена! Все данные сохранены в файл", fileName, "Введены следующие строки")
	fmt.Println(fullStr)
	file.WriteString(fullStr)
}