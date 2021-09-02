package main

import (
	"fmt"
	"math"
)

//Домашнее задание 7
func main() {

	fmt.Println("========================")
	fmt.Println("Задача 07.01.")
	fmt.Println("========================")
	task0701()
	fmt.Println("========================")
	fmt.Println("Задача 07.02.")
	fmt.Println("========================")
	task0702()
	fmt.Println("========================")
	fmt.Println("Задача 07.03.")
	fmt.Println("========================")
	task0703()
	fmt.Println("========================")
	fmt.Println("Задача 07.04.")
	fmt.Println("========================")
	task0704()

}
func task0701() {
	/*1. Зеркальные билеты 123 321
	Что нужно сделать
	Вывести сколько билетов находится среди всех шестизначных номеров от 100000 до 999999.

	Советы и рекомендации
	Для определения цифр числа, используйте цикл и операции целочисленного деления и получения остатка от целочисленного деления. При определенной реализации алгоритма,
	может быть полезно возведение 10 в степень, это можно сделать циклом, или воспользоваться готовой функцией: math.Pow10(n). Где в качестве n передается степень.
	Необходимо учитывать, что math.Pow возвращает дробное число, и чтобы деление было целочисленным необходимо воспользоваться конструкцией:
	var a int
	a = int(math.Pow10(3))
	В результате, в переменную а будет записана 1000.

	Что оценивается
	В коде программы должны быть два вложенных цикла, программа должна вывести 900 */

	fmt.Println("Зеркальные билеты")
	var countTicket int // Количество зеркальных билетов
	for ticketNumber := 100; ticketNumber <= 999; ticketNumber++ {

		num1 := ticketNumber / 100 % 10 //Сотни
		num2 := ticketNumber / 10 % 10  //Десятки
		num3 := ticketNumber % 10       //Еденицы

		num := num1*int(math.Pow10(5)) + num2*int(math.Pow10(4)) + num3*int(math.Pow10(3)) + num3*int(math.Pow10(2)) + num2*10 + num1

		fmt.Println(num, "num1 =", num1, "num2 =", num2, "num3 =", num3)
		countTicket++
	}
	fmt.Println("Количество зеркальных билетов:", countTicket)
}
func task0702() {
	/*2. Шахматная доска
	Что нужно сделать
	Одним из видов компьютерного искусства является псевдографика, когда из символов создаются картины.
	Предлагаю попробовать вывести два изображения.
	В рамках этой задачи, необходимо запросить у пользователя размер шахматной доски в клетках, и вывести шахматную доску на экран.
	Белые поля выводим звездочкой, а черные пробелом. Если пользователь введет 4, то мы должны увидеть:
	 * *
	* *
	 * *
	* *

	Советы и рекомендации
	Можно использовать fmt.Print и fmt.Println

	Что оценивается
	Правильность отображения доски для размеров 5 и 8 */
	fmt.Println("Шахматная доска")

	var sizeСhessboard int     // размер доски
	var cellWhite string = " " // беля клетка
	var cellBlack string = "*" // черная клетка

	fmt.Print("Введите размер доски:")
	fmt.Scan(&sizeСhessboard)

	for i := 0; i < sizeСhessboard; i++ {
		for j := 0; j < sizeСhessboard; j++ {

			if j%2 == 0 && i%2 == 0 {
				fmt.Print(cellWhite)
			} else if j%2 != 0 && i%2 == 0 {
				fmt.Print(cellBlack)
			} else if j%2 != 0 && i%2 != 0 {
				fmt.Print(cellWhite)
			} else if j%2 == 0 && i%2 != 0 {
				fmt.Print(cellBlack)
			}
		}
		fmt.Println()
	}
}
func task0703() {
	/*3. Вывод елочки
	Что нужно сделать

	Усложним задачу рисования, и попробуем вывести елочку.
	В первой строке необходимо вывести одну звездочку, во второй - на две больше, в третьей - еще на две больше и так до количества строк указанных пользователем.
	Например, если пользователь введет 5 то вывод должен иметь вид:
	    *
	   ***
	  *****
	 *******
	*********

	Советы и рекомендации
	Необходимо выводить пробелы и затем, звездочки.
	Посмотрите, как количество пробелов и звездочек, в каждой строке, связано с введенным количеством строк и номером текущей строки.
	Внутри цикла по строкам, скорее всего, понадобиться два цикла, один для вывода пробелов, второй для вывода звездочек.

	Что оценивается

	Правильность вывода елочки, что она симметрична и количество строк соответствует введенному пользователем. */
	fmt.Println("Вывод елочки")
	var val int // размер елочки

	fmt.Print("Введите размер елочки:")
	fmt.Scan(&val)

	for i := 0; i < val; i++ {
		for j := 0; j < (val*2)-1; j++ {
			if val-1-i <= j && val-1+i >= j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func task0704() {
	/*4. Сколько нужно купить билетов, для получения счастливого (дополнительно)
	Что нужно сделать
	Давайте поможем человеку, желающему ездить только со счастливым билетом (сумма первых трех цифр должна быть равна сумме последних трех цифр).
	Для этого он готов покупать любое количество билетов, но с одной стороны, он не хочет переплачивать, а с другой, он хочет быть твердо уверен,
	что среди купленных билетов будет счастливый. Необходимо написать программу, которая для билетов в диапазоне от 100000 до 999999 выведет,
	сколько минимум надо купить билетов, чтобы среди них оказался счастливый, при учете того, что номер текущего билета мы не знаем.
	Или, иными словами, нам надо найти максимальное расстояние между счастливыми билетами.
	Так после 100001 до следующего счастливого числа 100010 придется купить 9 билетов,
	а между 400220 и следующим счастливым билетом 400301 расстояние будет уже в 81 билет.

	Советы и рекомендации
	Необходимо запоминать предыдущий счастливый билет и максимальное расстояние найденное ранее.
	При нахождении очередного счастливого билета, необходимо находить расстояние до предыдущего и сравнивать его с максимальным.
	Если новое больше, то запоминать его. В любом случае, текущее счастливое число запоминать, как предыдущее и проверять числа дальше, до следующего счастливого.

	Что оценивается
	Программа должна вывести: 1001*/

	fmt.Println("Сколько нужно купить билетов, для получения счастливого (дополнительно)")

	var luckyTiketMax int          // последний счастливый билет
	var luckyTiketBefMax int       // предпоследний счастливый билет
	var distanceMax int            // максимальное растояние между билетами
	var distance int               // растояние между билетами
	var tiketLuckyOld int = 100000 // старый счастливый билет

	for ticketNumber := 100000; ticketNumber <= 999999; ticketNumber++ {

		num1 := ticketNumber / int(math.Pow10(5)) % 10 //Сотни тысяч
		num2 := ticketNumber / int(math.Pow10(4)) % 10 //Десятки тысяч
		num3 := ticketNumber / int(math.Pow10(3)) % 10 //Тысячи
		num4 := ticketNumber / int(math.Pow10(2)) % 10 //Сотни
		num5 := ticketNumber / 10 % 10                 //Десятки
		num6 := ticketNumber % 10                      //Еденицы

		if (num1 + num2 + num3) == (num4 + num5 + num6) {
			distance = ticketNumber - tiketLuckyOld
			if distance > distanceMax {
				luckyTiketBefMax = tiketLuckyOld
				luckyTiketMax = ticketNumber
				distanceMax = distance
			}
			tiketLuckyOld = ticketNumber
		}
	}
	fmt.Println("Максимальная дистанция:", distanceMax, "предпоследний счастливый билет:", luckyTiketBefMax, "последний счастливый билет:", luckyTiketMax)
}
