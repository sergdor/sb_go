package main

import "fmt"

func main() {
  fmt.Println("========================")
  fmt.Println("Задача 1.")
  fmt.Println("========================")
  Task1()
  fmt.Println("========================")
  fmt.Println("Задача 2.")
  fmt.Println("========================")
  Task2()
  fmt.Println("========================")
  fmt.Println("Задача 3.")
  fmt.Println("========================")
  Task3()
  fmt.Println("========================")
  fmt.Println("Задача 4.")
  fmt.Println("========================")
  Task4()
  fmt.Println("========================")
  fmt.Println("Задача 5.")
  fmt.Println("========================")
  Task5()
}
func Task1(){
//Задача 1. Конструктор квестов
  fmt.Println("Конструктор текстовых квестов")
  fmt.Println("Введите исходные данные")

  fmt.Print("Название планеты:")
  var planet string
  fmt.Scanln(&planet)
  
  fmt.Print("Название системы:")
  var star string
  fmt.Scanln(&star)
  
  fmt.Print("Имя рейнджера:")
  var ranger string
  fmt.Scanln(&ranger)
  
  fmt.Print("Количество вознаграждения:")
  var money int
  fmt.Scan(&money)
  
  fmt.Println("============================")

  fmt.Print("Как вам, ", ranger, ", известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты ", planet, " системы ", star, ".\n\n")
  
  fmt.Print("Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар планеты ", planet, ", впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!\n\n")
  
  fmt.Print(ranger, ", вы должны прибыть на планету ", planet, " системы ", star, " и помочь выполнить план призыва. Мы готовы выплатить вам премию в ", money," кредитов за эту маленькую услугу.\n")
}
func Task2() {
//Задача 2. Симулятор маршрутки
  fmt.Println("Симулятор маршрутки")
    
  var busStop1, busStop2, busStop3, busStop4 string //Остановка 1..4
  
  var passengersOut int //Пассажиров вошло
  var passengersIn int  //Пассажиров вышло
  passengersTotal := 0  //Всего пассажиров в маршрутке
  allPassengersIn :=0   //Все вошедшие пассажиры
  fmt.Println("Введите название остановок:")
 
  fmt.Print("Остановка 1:")
  fmt.Scanln(&busStop1)
  fmt.Print("Остановка 2:")
  fmt.Scanln(&busStop2)
  fmt.Print("Остановка 3:")
  fmt.Scanln(&busStop3)
  fmt.Print("Остановка 4:")
  fmt.Scanln(&busStop4)
 
  fmt.Println("\n", "-----------Едем---------", "\n")
  
  
  fmt.Print("Прибываем на остановку ", busStop1, ". В салоне пассажиров:", passengersTotal, "\n")
  
  fmt.Print("Сколько пассажиров вышло на остановке?:")
  fmt.Scan(&passengersOut)
  fmt.Print("Сколько пассажиров зашло на остановке?:")
  fmt.Scan(&passengersIn)
  
  passengersTotal += passengersIn - passengersOut 
  allPassengersIn += passengersIn
  
  fmt.Print("Отправляемся с остановки ", busStop1, ". В салоне пассажиров:", passengersTotal, "\n")
  fmt.Println("\n", "-----------Едем---------", "\n")
  
  
  fmt.Print("Прибываем на остановку ", busStop2, ". В салоне пассажиров:", passengersTotal, "\n")
  
  fmt.Print("Сколько пассажиров вышло на остановке?:")
  fmt.Scan(&passengersOut)
  fmt.Print("Сколько пассажиров зашло на остановке?:")
  fmt.Scan(&passengersIn)
  
  passengersTotal += passengersIn - passengersOut 
  allPassengersIn += passengersIn
  
  fmt.Print("Отправляемся с остановки ", busStop2, ". В салоне пассажиров:", passengersTotal, "\n")
  fmt.Println("\n", "-----------Едем---------", "\n")
  
  
  fmt.Print("Прибываем на остановку ", busStop3, ". В салоне пассажиров:", passengersTotal, "\n")
  
  fmt.Print("Сколько пассажиров вышло на остановке?:")
  fmt.Scan(&passengersOut)
  fmt.Print("Сколько пассажиров зашло на остановке?:")
  fmt.Scan(&passengersIn)
  
  passengersTotal += passengersIn - passengersOut 
  allPassengersIn += passengersIn
  
  fmt.Print("Отправляемся с остановки ", busStop3, ". В салоне пассажиров:", passengersTotal, "\n")
  fmt.Println("\n", "-----------Едем---------", "\n")
  
  
  fmt.Print("Прибываем на остановку ", busStop4, ". В салоне пассажиров:", passengersTotal, "\n")
  
  fmt.Print("Сколько пассажиров вышло на остановке?:")
  fmt.Scan(&passengersOut)
  fmt.Print("Сколько пассажиров зашло на остановке?:")
  fmt.Scan(&passengersIn)
  
  passengersTotal += passengersIn - passengersOut 
  allPassengersIn += passengersIn
  
  fmt.Print("Отправляемся с остановки ", busStop4, ". В салоне пассажиров:", passengersTotal, "\n")
  fmt.Println("\n", "-----------Едем---------", "\n")
  
  allMoney := allPassengersIn * 20 //Всего денег

  salary  := allMoney / 4 //Зарплата водителя
  fuel    := allMoney / 5 //Расходы на топливо
  taxes   := allMoney / 5 //Налоги
  repair  := allMoney / 5 //Расходы на ремонт машины
  
  fmt.Println("\n", "==========================", "\n")
  fmt.Println("\n", "Всего заработали:", allMoney ," руб.")
  fmt.Println("\n", "Зарплата водителя:", salary ," руб.")
  fmt.Println("\n", "Расходы на топливо:", fuel ," руб.")
  fmt.Println("\n", "Налоги:", taxes ," руб.")
  fmt.Println("\n", "Расходы на ремонт машины:", repair ," руб.")
  fmt.Println("\n", "Итого доход:", allMoney - salary - fuel - taxes - repair ," руб.")
}
func Task3(){
//Задача 3. Обмен местами
  fmt.Println("Обмен местами")
  
  a := 42
  b := 153
  
  fmt.Println("Исходные данные")
  fmt.Println("a:", a)
  fmt.Println("b:", b) 

  c := a
  a = b
  b = c

  fmt.Println("a:", a)
  fmt.Println("b:", b) 
}
func Task4(){
//Задача 4. Злостные вредители
  fmt.Println("Бамбук против вредителей")
	var startHeight, growth, losses, targetHeight, targetDays, days, height int

	fmt.Println("Введите параметры для расчета:")

	fmt.Print("Высота санженца:")
	fmt.Scan(&startHeight)
	fmt.Print("Скорость роста:")
	fmt.Scan(&growth)
	fmt.Print("Скорость его поедания гусеницами:")
	fmt.Scan(&losses)
	fmt.Print("Количество дней(для первой части задания):")
	fmt.Scan(&targetDays)

	fmt.Print("Целевая высота(см):")
	fmt.Scan(&targetHeight)

	height = (growth-losses)*(targetDays-1) + growth/2 + startHeight

	fmt.Println("Высота бамбука к середине ", targetDays, " дня:", height)
    
	days = (targetHeight - startHeight - growth) / (growth - losses) + 2 //почему 2 скорее всего связанно с округлением

	fmt.Print("Полных дней до целевой высоты ", +targetHeight, " cм.: ", days, " дней \n")
}
func Task5(){
//Задача 5. Повышаем градус сложности
  fmt.Println("Повышаем градус сложности")
  
  a := 42
  b := 153
  fmt.Println("Исходные данные")
  fmt.Println("a:", a,"b:", b)
  
  a = a + b
  b = a - b
  a = a - b

  fmt.Println("a:", a,"b:", b)
  /*
  Проблемы возникнут если сумма 2 чисел привышает максимально возможное число для типа integer 
  Преимущества в отсутствии необходимости добавлять буферную переменную не забивая память, но увеличивает нагрузку на процессор операциями сложения и вычитания
  */
}