package main

import (
	"fmt"
	"math"
)

//Домашнее задание 4
func main() {
	fmt.Println("========================")
	fmt.Println("Задача 05.01.")
	fmt.Println("========================")
	task0501()
	fmt.Println("========================")
	fmt.Println("Задача 05.02.")
	fmt.Println("========================")
	task0502()
	fmt.Println("========================")
	fmt.Println("Задача 05.03.")
	fmt.Println("========================")
	task0503()
	fmt.Println("========================")
	fmt.Println("Задача 05.04.")
	fmt.Println("========================")
	task0504()
	fmt.Println("========================")
	fmt.Println("Задача 05.05.")
	fmt.Println("========================")
	task0505()
	fmt.Println("========================")
	fmt.Println("Задача 05.06.")
	fmt.Println("========================")
	task0506()
	fmt.Println("========================")
	fmt.Println("Задача 05.07.")
	fmt.Println("========================")
	task0507()
	fmt.Println("========================")
	fmt.Println("Задача 05.08.")
	fmt.Println("========================")
	task0508()
}
func task0501() {
	/*Задача 1. Определение координатной плоскости точки
	Что нужно сделать
	В различных программах часто приходится работать с координатами, это могут быть системы проектирования или игры.
	Ключевой момент в таких программах, это работа с системой координат. Давайте поможем пользователю определить, к какой координатной четверти принадлежит точка.
	Пользователь вводит числа x, y, а программе необходимо вывести, какой координатной четверти принадлежит точка с координатами (x, y), при условии:
	    если обе координаты положительны, то точка находится в первой четверти координатной плоскости;
	    если координата х отрицательна, а координата у положительна, то точка находится во второй четверти;
	    если обе координаты отрицательны, то число находится в третьей четверти;
	    если координата х положительна, а координата у отрицательна, то точка лежит в четвертой четверти.

	Советы и рекомендации
	Использовать логическое умножение(И).*/
	fmt.Println("Определение координатной плоскости точки")
	var x int                                     // координата Х
	var y int                                     // координата Y
	var numberQuarters int                        // номер четверти
	msgErrInput := "Не верно введенна координата" // сообщение об ошибке при неверном вводе координат
	for x == 0 {
		fmt.Print("Введите координату X:")
		fmt.Scan(&x)
		if x == 0 {
			fmt.Println(msgErrInput)
		}
	}
	for y == 0 {
		fmt.Print("Введите координату Y:")
		fmt.Scan(&y)
		if y == 0 {
			fmt.Println(msgErrInput)
		}
	}

	if x > 0 && y > 0 {
		numberQuarters = 1
	} else if x < 0 && y > 0 {
		numberQuarters = 2
	} else if x < 0 && y < 0 {
		numberQuarters = 3
	} else if x > 0 && y < 0 {
		numberQuarters = 4
	}
	fmt.Println("Точка лежит в", numberQuarters, "четверти")

}
func task0502() {
	/*Задача 2. Проверить, что одно из чисел положительное
	Что нужно сделать
	Проверка пользовательского ввода на различные ограничения является частой задачей.
	Давайте попросим пользователя ввести 3 числа и проверим, что хотя бы одно является положительным.
	Результат проверки необходимо сообщить пользователю.

	Советы и рекомендации
	Использовать логическое сложение(ИЛИ).*/
	fmt.Println("Проверить, что одно из чисел положительное")
	var val1, val2, val3 int // пользовательский ввод
	fmt.Print("Введите первое число:")
	fmt.Scan(&val1)
	fmt.Print("Введите второе число:")
	fmt.Scan(&val2)
	fmt.Print("Введите третье число:")
	fmt.Scan(&val3)
	if val1 > 0 || val2 > 0 || val3 > 0 {
		fmt.Println("В веденных данных есть положительное число")
	} else {
		fmt.Println("В веденных данных нет положительного числа")
	}

}
func task0503() {
	/*Задача 3. Проверить, что есть совпадающие числа
	Что нужно сделать
	Пользователь при вводе может ошибиться, и многие программы пытаются привлечь внимание к таким ошибкам.
	Например, текстовые редакторы подсвечивают ошибки в словах, электронные таблицы — в формулах.
	Давайте реализуем программу, запрашивающую у пользователя три числа, и выводящую подсказку, если два или более числа совпадают.

	Советы и рекомендации
	Использовать логическое сложение(ИЛИ).*/
	fmt.Println("Проверить, что есть совпадающие числа")
	var val1, val2, val3 int // пользовательский ввод
	fmt.Print("Введите первое число:")
	fmt.Scan(&val1)
	fmt.Print("Введите второе число:")
	fmt.Scan(&val2)
	fmt.Print("Введите третье число:")
	fmt.Scan(&val3)
	if val1 == val2 || val1 == val3 || val2 == val3 {
		fmt.Println("Есть совпадающие числа")
	} else {
		fmt.Println("Нет совпадающих чисел")
	}

}
func task0504() {
	/*Задача 4. Сумма без сдачи
	Что нужно сделать
	Программное обеспечение банкоматов постоянно решает задачу, как имеющимися купюрами сформировать сумму, введенную пользователем.
	Давайте попробуем решить похожую задачу, и определить, сможет ли пользователь заплатить за товар без сдачи или нет.
	Для этого он будет вводить стоимость товара и номиналы трех монет.
	Примеры работы программы:
	сумма:
	7
	монеты:
	1
	2
	5
	Пользователь может оплатить без сдачи
	сумма:
	9
	монеты:
	5
	2
	5
	Пользователь не может оплатить без сдачи

	Советы и рекомендации
	Необходимо проверить все сценарии, когда сумма может быть сформирована одной монетой, двумя или всеми тремя.*/
	fmt.Println("Сумма без сдачи")

	var coin1, coin2, coin3 int // 3 монеты
	var valueGoods int          // стоимость товара

	fmt.Print("Введите стоимость товара:")
	fmt.Scan(&valueGoods)

	fmt.Print("Введите номенал первой монеты:")
	fmt.Scan(&coin1)
	fmt.Print("Введите номенал второй монеты:")
	fmt.Scan(&coin2)
	fmt.Print("Введите номенал третей монеты:")
	fmt.Scan(&coin3)
	if coin1+coin2+coin3 >= valueGoods {
		if coin1+coin2+coin3 == valueGoods || coin1+coin2 == valueGoods || coin1+coin3 == valueGoods || coin2+coin3 == valueGoods || coin1 == valueGoods || coin2 == valueGoods || coin3 == valueGoods {
			fmt.Println("Пользователь может оплатить без сдачи")
		} else {
			fmt.Println("Пользователь не может оплатить без сдачи")
		}
	} else {
		fmt.Println("Недостаточно средств!")
	}
}

func task0505() {
	/*Задача 5. Определение максимальных процентов
	Что нужно сделать
	Задача учебная, и человек с ней справится сам, но давайте научим нашу программу определять, какие две из предложенных ставок по депозитам являются максимальными.
	Пользователь будет вводить три процентные ставки, а программа должна вывести, какие две из них являются более выгодными.

	Советы и рекомендации
	Задачу можно решать несколькими способами, например, от противного.	*/
	fmt.Println("Определение максимальных процентов")

	var depositRate1, depositRate2, depositRate3 int // Ставки по депозитам

	fmt.Print("Введите ставку по первому депозиту:")
	fmt.Scan(&depositRate1)
	fmt.Print("Введите ставку по второму депозиту:")
	fmt.Scan(&depositRate2)
	fmt.Print("Введите ставку по третьему депозиту:")
	fmt.Scan(&depositRate3)

	if depositRate1 == depositRate2 && depositRate1 == depositRate3 {
		fmt.Println("Ставки равны")
	} else {
		if depositRate1 >= depositRate2 || depositRate1 >= depositRate3 { //ставка 1 является одной из максимальных
			fmt.Println("Ставка 1 является выгодной", depositRate1)
			if depositRate2 >= depositRate3 { //ставка 2 является одной из максимальных
				fmt.Println("Ставка 2 является выгодной", depositRate2)
			} else { //ставка 3 является одной из максимальных
				fmt.Println("Ставка 3 является выгодной", depositRate3)
			}
		} else { // ставка 2 и 3 являются максимальными
			fmt.Println("Ставка 2 является выгодной", depositRate2)
			fmt.Println("Ставка 3 является выгодной", depositRate3)
		}
	}
}

func task0506() {
	/*Задача 6. Решение квадратного уравнения
	Что нужно сделать
	Что компьютеры делают быстрее людей? Проводят вычисления!
	Давайте напишем программу, решающую квадратные уравнения. Пользователь вводит коэффициенты a, b и c квадратного уравнения.
	Программа должна вывести корни уравнения. Для решения уравнения необходимо:
	Вычислить дискриминант по формуле:
	Если D>0, будет два различных корня, которые находятся по формуле.
	Если D = 0, будет один корень, который находится по формуле.
	Если D < 0, то вывести, что корней нет.

	Советы и рекомендации
	Для возведения b в квадрат, можно воспользоваться умножением, а можно функцией math.Pow (при ее использовании переменная b должна быть типа float64).
	Для вычисления квадратного корня необходимо воспользоваться функцией math.Sqrt.*/

	fmt.Println("Решение квадратного уравнения") //x = (-b±✓D)/2a, где D = b²-4ac

	var a, b, c float64

	fmt.Print("Введите a:")
	fmt.Scan(&a)
	fmt.Print("Введите b:")
	fmt.Scan(&b)
	fmt.Print("Введите c:")
	fmt.Scan(&c)
	d := math.Pow(b, 2) - 4*a*c
	fmt.Println("d =", d)
	//d = 0
	if d < 0 {
		fmt.Println("корней нет")
	} else if d == 0 {
		fmt.Println("корень один")
		fmt.Println(-b / (2 * a))
	} else if d > 0 {
		fmt.Println("корней два")
		fmt.Println((-b + math.Sqrt(d)) / 2 * a)
		fmt.Println((-b - math.Sqrt(d)) / 2 * a)
	}
}
func task0507() {
	/*Задача 7. Счастливый билет
	Что нужно сделать
	Все мы в детстве, да и не только в детстве, получив бумажный билет, проверяли, а не является ли он “счастливым”.
	Давайте напишем программу, в которую пользователь будет вводить четырехзначный номер билета, а программа будет выводить,
	является ли он счастливым (сумма старших двух разрядов равна сумме двух младших разрядов) или даже зеркальным.
		Например:
		1234 -> обычный билет
		3425 -> счастливый билет
		1221 -> зеркальный билет

	Советы и рекомендации
	При решении задачи необходимо применить целочисленное деление и остаток от деления.	*/
	fmt.Print("Счастливый билет")
	var ticketNumber int
	fmt.Print("Введите номер билета:")
	fmt.Scan(&ticketNumber)

	num1 := ticketNumber / 1000 % 10 //Тысячи
	num2 := ticketNumber / 100 % 10  //Сотни
	num3 := ticketNumber / 10 % 10   //Десятки
	num4 := ticketNumber % 10        //Еденицы
	fmt.Print(ticketNumber, " -> ")
	if num1 == num4 && num2 == num3 {
		fmt.Println("зеркальный билет")
	} else if num1+num2 == num3+num4 {
		fmt.Println("счастливый билет")
	} else {
		fmt.Println("обычный билет")
	}
}
func task0508() {
	/*Задача 8. Игра “Угадывание числа” (дополнительно)
	Что нужно сделать
	Ну, и какой же компьютер без игр? Давайте научим его играть в “угадывание чисел”. Пользователь загадывает число от 1 до 10 (включительно).
	Программа пытается это число угадать, для этого она выводит число, а пользователь должен ответить угадала программа, загаданное число больше или меньше.

	Советы и рекомендации
	Программа не должна делать больше 4 попыток в процессе угадывания.
	*/
	fmt.Println("Игра 'Угадывание числа'")
	fmt.Println(" > загаданное число больше")
	fmt.Println(" < загаданное число меньше")
	fmt.Println(" = программа отгадала число")
	fmt.Println("---------------------------")

	fmt.Println("Загадайте число от 1 до 10(включительно):")
	var userInput string
	val := 5
	min := 1
	max := 11
	var tryNumber int

	for userInput != "=" {
		tryNumber++
		val = (min + max) / 2
		fmt.Print("Это число ", val, "?\n")
		fmt.Scan(&userInput)
		if userInput == "=" {
			fmt.Print("Число угаданно с ", tryNumber, " попытки")
		} else if userInput == ">" {
			min = val
		} else if userInput == "<" {
			max = val
		} else {
			fmt.Println("Ошибка ввода введите '>', '<' или '='")
			tryNumber--
		}

		if tryNumber == 4 && userInput != "=" {
			fmt.Println("Программе не удалось угадать число с 4х попыток")
			break
		}
	}
}
