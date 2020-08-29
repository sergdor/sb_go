package main

import "fmt"

//Домашнее задание 4
func main() {
	fmt.Println("========================")
	fmt.Println("Задача 1.")
	fmt.Println("========================")
	//task1()
	fmt.Println("========================")
	fmt.Println("Задача 2.")
	fmt.Println("========================")
	//task2()
	fmt.Println("========================")
	fmt.Println("Задача 3.")
	fmt.Println("========================")
	//task3()
	fmt.Println("========================")
	fmt.Println("Задача 4.")
	fmt.Println("========================")
	//task4()
	fmt.Println("========================")
	fmt.Println("Задача 5.")
	fmt.Println("========================")
	//task5()
	fmt.Println("========================")
	fmt.Println("Задача 6.")
	fmt.Println("========================")
	//task6()
	fmt.Println("========================")
	fmt.Println("Задача 7.")
	fmt.Println("========================")
	task7()

}

func task1() {
	/*Задача 1. Баллы ЕГЭ
	  Сумма проходных баллов в институт составляет 275 баллов.
	  Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет, поступил ли он в институт или нет.
	*/

	var val, amountExams int
	messageBal := "Количество баллов за экзамен"
	fmt.Println("Баллы ЕГЭ \n")
	fmt.Println("Введие параметры:")

	fmt.Print(messageBal, "1:")
	fmt.Scan(&val)
	amountExams += val
	fmt.Print(messageBal, "2:")
	fmt.Scan(&val)
	amountExams += val
	fmt.Print(messageBal, "3:")
	fmt.Scan(&val)
	amountExams += val

	if amountExams >= 275 {
		fmt.Println("Поздравляю! Вы приняты в институт!")
	} else {
		fmt.Println("Вы не приняты в институт!")
	}
}

func task2() {
	/*Задача 2. Три числа
	  Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
	*/
	var val1, val2, val3 int
	fmt.Println("Три числа \n")
	fmt.Print("Введите первое число:")
	fmt.Scan(&val1)
	fmt.Print("Введите второе число:")
	fmt.Scan(&val2)
	fmt.Print("Введите третье число:")
	fmt.Scan(&val3)

	if val1 > 5 {
		fmt.Println("Первое число", val1, "больше 5")
	}
	if val2 > 5 {
		fmt.Println("Второе число", val2, "больше 5")
	}
	if val3 > 5 {
		fmt.Println("Теретье число", val3, "больше 5")
	}
}
func task3() {
	/*Задача 3. Банкомат
	  Пользователи банкомата хотят снимать деньги. Но банкомат умеет выдавать только купюры по 100 рублей, а максимальная сумма снятия — 100 000 рублей.
	  Напишите программу, которая проверяет допустимость суммы средств, которую ввёл пользователь. Сумма не должна быть меньше 100 и больше 100 000 руб.
	*/
	var val int
	fmt.Println("Банкомат \n")
	fmt.Print("Введите сумму для снятия:")
	fmt.Scan(&val)

	if val < 100 {
		fmt.Println("Банкомат не может выдать сумму меньше 100 руб.")
	} else if val > 100000 {
		fmt.Println("В банкомате не достаточно средств")
	} else if val%100 == 0 {
		fmt.Println("Выдача наличных", val, " руб.")
	} else {
		fmt.Println("Банкомат может выдать только ратное 100 руб.")
	}
}

func task4() {
	/*Задача 4. Три числа: еще попытка
	  Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
	*/var val1, val2, val3, numberHigher5 int
	fmt.Println("Три числа: еще попытка \n")
	fmt.Print("Введите первое число:")
	fmt.Scan(&val1)
	fmt.Print("Введите второе число:")
	fmt.Scan(&val2)
	fmt.Print("Введите третье число:")
	fmt.Scan(&val3)

	if val1 > 5 {
		numberHigher5++
	}
	if val2 > 5 {
		numberHigher5++
	}
	if val3 > 5 {
		numberHigher5++
	}

	fmt.Println("Чесел больше 5:", numberHigher5)
}

func task5() {
	/*Задача 5. Ресторан
	  В ресторане действуют следующие правила:
	        по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый!
	        по пятницам, если сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%;
	        если число гостей в одной компании превышает 5 человек, автоматически включается надбавка на обслуживание 10%.
	    Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
	*/
	var guests, check, discount, surcharge int
	var day string
	fmt.Println("Ресторан \n")
	fmt.Print("Введите день недели:")
	fmt.Scan(&day)
	fmt.Print("Введите количество гостей:")
	fmt.Scan(&guests)
	fmt.Print("Введите сумму чека:")
	fmt.Scan(&check)

	if day == "Понедельник" {
		discount = check * 10 / 100
	} else if day == "Пятница" {
		if check > 10000 {
			discount = check * 5 / 100
		}
	}
	if guests > 5 {
		surcharge = check * 10 / 100
	}
	fmt.Print("Сумма к оплате:", check-discount+surcharge, "\n")
}

func task6() {
	/*
	   Задача 6. Студенты* (сложная задача)
	   Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов, на K групп.
	   Напишите программу, которая поможет старосте сделать это:
	   он вводит N, K и порядковый номер студента, а программа определяет, в какую группу он попадает.
	   Подсказка: в одну группу могут попадать студенты из разных частей списка.
	*/
	var studentsNumber int     // Количество студентов
	var groupsNumber int       // Количество групп
	var studentsMaxInGroup int // Максимально студентов вгруппе
	var studentNumber int      // Номер текущего студента
	var studentsInGroup int    // Студентов в группе
	var groupNumber int        // Группа номер

	fmt.Println("Студенты* (сложная задача)")
	fmt.Print("Введите количество студентов:")
	fmt.Scan(&studentsNumber)
	fmt.Print("Введите количество групп:")
	fmt.Scan(&groupsNumber)
	studentsMaxInGroup = studentsNumber / groupsNumber
	if studentsNumber/groupsNumber > studentsMaxInGroup {
		studentsMaxInGroup++
	}
	fmt.Print("Мксимальное количество студентов в группе:", studentsMaxInGroup, "\n")
	for i := 1; i <= studentsNumber; i++ {
		if studentsInGroup == studentsMaxInGroup || groupNumber == 0 {
			studentsInGroup = 0
			groupNumber++
			fmt.Print("Заполнение группы №", groupNumber, "\n")
		}
		fmt.Print("Введите порядковый номер студента:")
		fmt.Scan(&studentNumber)
		studentsInGroup++
	}
}
func task7() {
	/*
		Задача 7. Прогрессивный налог* (сложная задача)
		В последнем уроке мы писали программу, вычисляющую сумму налога по прогрессивной шкале в зависимости от полученного
		заработка: 13% на доход до 10000 руб., 20% на доход от 10000 до 50000 руб., 30% на доход выше 50000 руб.
		Однако, во многих странах, использующих такую шкалу, эта сумма вычисляется более сложным способом.
		А именно, налоговая ставка 30% на доход выше 50000 руб. означает, что 30% уплачивается не со всей суммы,
		а лишь с той ее части, которая превосходит 50000 руб. Аналогично, ставка 20% на доход от 10000 до 50000 руб.
		обязывает уплатить 20% лишь с той части суммы, которая превосходит 10000 руб., но не превосходит 50000 руб.
		Так, например, с дохода 100000 руб. придется заплатить такой налог:
		30% * (100000-50000) + 20% * (50000-20000) + 13% * 10000 = 15000+6000+1300=22300.
		А с дохода 30000 — такой: 20% * (30000-20000) + 13% * 10000 = 2000+1300=3300.
		Напишите программу, которая спрашивает у пользователя его доход и высчитывает сумму налога для него по вышеописанным правилам.
	*/

<<<<<<< HEAD
	fmt.Println("Прогрессивный налог* (сложная задача)")
	var profit int               // Прибыль
	var tax1, tax2, tax3 float64 // Сумма налага 1,2,3
	var summTax1 int             // 1 сумма с которой начинается налог
	summTax2 := 10000            // 2 сумма с которой начинается налог
	summTax3 := 50000            // 3 сумма с которой начинается налог
	taxPercent1 := 0.13          // 1 процент налога
	taxPercent2 := 0.20          // 2 процент налога
	taxPercent3 := 0.30          // 3 процент налога
	var tax float64              // Общая сумма налога
	for profit == 0 {
		fmt.Print("Введите сумму прибыли:")
		fmt.Scan(&profit)
		if profit == 0 {
			fmt.Println("Некорректно введены данные! Повторите ввод данных!")
		}
	}

	if profit >= summTax3 {
		tax3 = float64(profit-summTax3) * taxPercent3
		tax2 = float64(summTax3) * taxPercent2
		tax1 = float64(summTax2) * taxPercent1
		tax = tax1 + tax2 + tax3
		//tax = (float64(summTax2) * taxPercent1) + float64(summTax3) * taxPercent2 + float64(profit-summTax3) * taxPercent3
	} else if profit >= summTax2 {
		tax2 = float64(profit-summTax2) * taxPercent2
		tax1 = float64(summTax2) * taxPercent1
		tax = tax1 + tax2
		//tax = (float64(summTax2) * taxPercent1) + float64(summTax3) * taxPercent2
	} else if profit >= summTax1 {
		tax1 = float64(profit-summTax1) * taxPercent1
		tax = tax1
		//tax = float64(profit-summTax1) * taxPercent1
	}
	fmt.Print("Сумма налога ", taxPercent1*100, "%:", tax1, "\n")
	if tax2 > 0 {
		fmt.Print("Сумма налога ", taxPercent2*100, "%:", tax2, "\n")
	}
	if tax3 > 0 {
		fmt.Print("Сумма налога ", taxPercent3*100, "%:", tax3, "\n")
	}

	fmt.Print("Общая сумма налога составила:", tax, "\n")
=======
>>>>>>> ec0458d624c16524c9191fe4e2748642a7ba400f
}
