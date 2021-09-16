package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("========================")
	fmt.Println("Задача 12.02.")
	fmt.Println("========================")
	task1202()
	fmt.Println("========================")
	fmt.Println("Задача 12.03.")
	fmt.Println("========================")
	task1203()
	fmt.Println("========================")
	fmt.Println("Задача 12.04.")
	fmt.Println("========================")
	task1204()
	fmt.Println("========================")
	fmt.Println("Задача 12.05.")
	fmt.Println("========================")
	task1205()
	fmt.Println("========================")
	fmt.Println("Задача 12.05.")
	fmt.Println("========================")
	task1206()
}

func task1202() {
	//Урок №2 Работа с файлами
	//Написать программу, которая на вход получала бы строку, введенную пользователем, а в файл писала: № строки,
	//дата и сообщение в формате:
	//2020-02-10 15:00:00 продам гараж
	//при вводе слова “выход” производился бы выход из программы
	fileName := "hw12.2.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Ошибка создания файла", err)
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
	_, err = file.WriteString(fullStr)
	if err != nil {
		fmt.Print("Ошибка записи в файл", err)
		return
	}
	fmt.Println("Работа программы завершена! Все данные сохранены в файл", fileName, "Введены следующие строки")
	fmt.Println(fullStr)
}

func task1203() {
	//Урок №3 интерфейс io.Reader
	//Написать программу, которая полностью вычитывала бы файл из предыдущего домашнего задания без использования ioutil
	//Подсказка: для получения размера файла у файла есть метод Stat(), который возвращает информацию о файле и ошибку.

	fileName := "hw12.2.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле", err)
		return
	}

	buf := make([]byte, stat.Size())
	if _, err := io.ReadFull(file, buf); err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return
	}
	fmt.Println("Содержимое файла", fileName)
	fmt.Printf("%s", buf)
}

func task1204() {
	//Урок №4 уровни доступа
	//Создать файл только для чтения и проверить, что в него нельзя записать данные

	fileName := "hw12.4.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Ошибка создания файла", err)
		return
	}
	file.Close()

	if err = os.Chmod(fileName, 0444); err != nil {
		fmt.Println("Ошибка установки прав", err)
		return

	}

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Print("Ошибка открытия файла", err)
		return
	}

	if _, err = file.WriteString("Test"); err != nil {
		fmt.Println("Ошибка записи в файл", err)
		return
	}
}

func task1205() {
	//Урок №6 пакет ioutil
	//переписать задачи из урока 2 и 3 на пакет ioutil

	//Урок №2 Работа с файлами
	//Написать программу, которая на вход получала бы строку, введенную пользователем, а в файл писала: № строки,
	//дата и сообщение в формате:
	//2020-02-10 15:00:00 продам гараж
	//при вводе слова “выход” производился бы выход из программы
	fileName := "hw12.5.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Ошибка создания файла", err)
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
	buf := []byte(fullStr)
	err = ioutil.WriteFile(fileName, buf, 0666)
	if err != nil {
		fmt.Print("Ошибка записи в файл", err)
		return
	}

	fmt.Println("Работа программы завершена! Все данные сохранены в файл", fileName, "Введены следующие строки")
	fmt.Println(fullStr)
	//file.WriteString(fullStr)

	//Урок №3 интерфейс io.Reader
	//Написать программу, которая полностью вычитывала бы файл из предыдущего домашнего задания без использования ioutil
	//Подсказка: для получения размера файла у файла есть метод Stat(), который возвращает информацию о файле и ошибку.

	fileName = "hw12.5.txt"

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	stat, err := file.Stat()

	file.Close()

	if err != nil {
		fmt.Println("Ошибка получения информации о файле", err)
		return
	}

	buf = make([]byte, stat.Size())

	buf, err = ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return
	}
	fmt.Println("Содержимое файла", fileName)
	fmt.Printf("%s", buf)

}

func task1206() {
	//Дополнительно задание со звездочкой*
	//Написать программу, которая на вход принимала бы интовое число
	//и для него генерировала бы все возможные комбинации круглых скобок, т.е.
	//на вход приходит число 3
	//на выходе:
	//["((()))","(()())","(())()","()(())","()()()"]
	//на вход 1
	//на выходе:
	//["()"]

	var val int
	fmt.Print("Введите число:")
	fmt.Scan(&val)
	//val = 3
	res := make([]string, 0)

	task1206_1(val, val, "", &res)

	fmt.Println(res)
}

func task1206_1(open, closing int, queue string, result *[]string) {
	if open == 0 {
		*result = append(*result, queue+strings.Repeat(")", closing))
		return
	}
	if closing > open {
		task1206_1(open, closing-1, queue+")", result)
	}
	task1206_1(open-1, closing, queue+"(", result)
}
