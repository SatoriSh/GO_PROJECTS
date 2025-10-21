package main

import (
	"fmt"
)

type (
	mile      int
	kilometer int
)

func main() {
	// =====================
	// |    типы данных    |
	// =====================
	fmt.Println("\n=======   типы данных   =======")
	{
		/*
			int8: представляет целое число от -128 до 127 и занимает в памяти 1 байт (8 бит)
			int16: представляет целое число от -32768 до 32767 и занимает в памяти 2 байта (16 бит)
			int32: представляет целое число от -2147483648 до 2147483647 и занимает 4 байта (32 бита)
			int64: представляет целое число от –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807 и занимает 8 байт (64 бита)
			uint8: представляет целое число от 0 до 255 и занимает 1 байт
			uint16: представляет целое число от 0 до 65535 и занимает 2 байта
			uint32: представляет целое число от 0 до 4294967295 и занимает 4 байта
			uint64: представляет целое число от 0 до 18 446 744 073 709 551 615 и занимает 8 байт
			byte: синоним типа uint8, представляет целое число от 0 до 255 и занимает 1 байт
			rune: синоним типа int32, представляет целое число от -2147483648 до 2147483647 и занимает 4 байта
			int: представляет целое число со знаком, которое в зависимости о платформы может занимать либо 4 байта, либо 8 байт. То есть соответствовать либо int32, либо int64.
			uint: представляет целое беззнаковое число только без знака, которое, аналогично типу int, в зависимости о платформы может занимать либо 4 байта, либо 8 байт. То есть соответствовать либо uint32, либо uint64.
		*/

		var varValueInt = 0
		var varValueString = "lol"
		fmt.Println("\nvar int value:", varValueInt, "\nvar string value:", varValueString)

		var valueInt = 1        // OR valueInt := int(1)
		var valueString = "LOL" // OR valueString := string("LOL")
		fmt.Println("\nint value:", valueInt, "\nstring value:", valueString)

		// можно создать одновременно несколько переменных
		var a, b, c int = 1, 2, 3
		println(a, b, c)

		// или обернуть в скобки
		var (
			varA int = 11
			varB int = 22
			varC int = 33
		)
		println(varA, varB, varC)

		const constValue float64 = 123.123

		const (
			pi float64 = 3.1415
			e  float64 = 2.7182
		)
		fmt.Println("\nconst float value:", constValue, "pi:", pi, "e:", e)

		// iota-идентификатор
		const // iota сбрасывается в 0
		(
			C0 = iota // здесь iota равно 0, увеличивается с каждой строкой
			C1        // увеличение на 1, iota равна 1
			C2 = iota // iota равна 2
		)

		fmt.Println("\nC0:", C0) // C0: 0
		fmt.Println("C1:", C1)   // C1: 1
		fmt.Println("C2:", C2)   // C2: 2

		const //  iota сбрасывается в 0
		(
			C3 = iota // С3 = 0
		)

		fmt.Println("C3:", C3) // C3: 0

	}

	// ==============================
	// |    поразрядные операции    |
	// ==============================
	fmt.Println("\n=======   поразрядные операции   =======")
	{
		// операции сдвига
		var a int = 2 << 2  // 2 в десятичной это 10, свдигаем его на 2 разряда влево получается 1000, 1000 в десятичной это 8
		var b int = 16 >> 3 // 16 в десятичной это 10000, на 3 разряда вправо получается 10, 10 в десятичной это 2

		fmt.Println("\n2 << 2:", a)
		fmt.Println("16 >> 3:", b)

		// форматирование в консоли
		var c int = 10
		fmt.Printf("c = %b\n", c)
		var cc int = c >> 2
		fmt.Printf("\n%b >> 2 = %b | %b - %d", c, cc, cc, cc)

		// логические операции
		/*
			&: И - оба включены
			|: ИЛИ - хотя бы один включен
			^: HOR - только один включен
			&^: AND NOT - включен только в первом
		*/
		var e int = 5 | 2  // 101 | 010 = 111 - 7 (если хотя бы один разряд равен 1, то сумма обоих разрядов равна 1)
		var f int = 6 & 2  // 110 & 010 = 010 - 2 (произведение обоих разрядов равно 1, если оба этих разряда равны 1)
		var g int = 5 ^ 1  // 101 : 001 = 100 - 4 (если только один разряд равен 1, то возвразаем 1)
		var h int = 5 &^ 6 // 101 &^ 110 = 001 - 1 (если в первом разряд равен 1, а во втором 0, возвращаем 1)

		fmt.Printf("\n\ne: %b | %b = %b\n", 5, 2, e)
		fmt.Printf("\nf: %b & %b = %b\n", 6, 2, f)
		fmt.Printf("\ng: %b ^ %b = %b\n", 5, 1, g)
		fmt.Printf("\nh: %b &^ %b = %b\n", 5, 6, h)
	}

	// =================
	// |    массивы    |
	// =================
	fmt.Println("\n=======   массивы   =======")
	{
		var numbers [5]int
		for i := 0; i < len(numbers); i++ {
			if i == len(numbers)-1 || i == 0 {
				fmt.Print("\n")
			}
			fmt.Print(numbers[i], " ") // 0 0 0 0 0
		}
		// или можно просто fmt.Println(numbers)

		var nums [3]int = [3]int{1, 2, 3}
		fmt.Println(nums)

		var names [2]string = [2]string{"Tom", "Peter"}

		names[0] += " - 1 name |"
		names[1] += " - 2 name"
		fmt.Println(names)
	}

	// ==============================
	// |    условные конструкции    |
	// ==============================
	fmt.Println("\n=======   условные конструкции   =======")
	{
		a := 5
		b := 4
		if a < b {
			fmt.Println("\na is less than b")
		} else if a > b {
			fmt.Println("\nb is less than a")
		} else {
			fmt.Println("\na is equal b")
		}

		switch {
		case a < b:
			fmt.Println("\na is less than b")
		case a > b:
			fmt.Println("\nb is less than a")
		default:
			fmt.Println("\na is equal b")
		}

		var c int = 7
		switch c + 8 {
		case 15:
			fmt.Println("c = 15")
		default:
			fmt.Println("?")
		}
	}

	// ===============
	// |    циклы    |
	// ===============
	fmt.Println("\n=======   циклы   =======")
	{
		var a int = 0
		for i := 0; i < 10; i++ {
			a += 10
			fmt.Println(a)
		}

		for i := 1; i < 10; i++ {
			for j := 1; j < 10; j++ {
				fmt.Print(i*j, "\t")
			}
			fmt.Println()
		}
		fmt.Println()

		var str string = "string"
		for index, value := range str {
			fmt.Printf("index: %d, value: %c\n\n", index, value)
		}

		var users [3]string = [3]string{"Tom", "Peter", "Garry"}
		for index, value := range users {
			fmt.Println(index, value)
		}

		var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
		var sum = 0

		for _, value := range numbers { // если index не нужен можно оставить прочерк
			if value < 0 {
				continue // переходим к следующей итерации
			}
			sum += value
		}
		fmt.Println("Sum:", sum) // Sum: 27

		var numberss = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var summ = 0

		for _, value := range numberss {
			if value > 4 {
				break // если число больше 4 выходим из цикла
			}
			summ += value
		}
		fmt.Println("Sum:", summ) // Sum: 10
	}

	// =================
	// |    функции    |
	// =================
	fmt.Println("\n=======   функции   =======")
	{
		add(10, 55)
		addMoreValues(1, 1, 1) // неопределенное количество параметров
		addMoreValues(13, 10)
		addMoreValues(7, 1, 1, 3, 0, 0, 4, 0)

		var numsArray []int = []int{10, 20, -1}
		fmt.Println(getMinValue(numsArray))
		numsArray = []int{1012, 20, -100, -120, 12389798, 0}
		fmt.Println(getMinValue(numsArray))
		numsArray = []int{100, 0, 1, 5, 0, 0}
		fmt.Println(getMinValue(numsArray))
	}

	// ===============================================================
	// |    тип функции, функция как параметр и результат функции    |
	// ===============================================================
	fmt.Println("\n=======   тип функции, функция как параметр и результат функции   =======")
	{
		fmt.Println()

		// переменная может быть функцией
		var a func(int, int) int = addAndReturn
		fmt.Println(a(2, 3)) // 5

		var f func(int, int) int = multiplyAndReturn
		fmt.Println("f - multiply", f(2, 3))
		f = addAndReturn // теперь переменная f указывает на функцию addAndReturn
		fmt.Println("f - add", f(2, 3))

		var result func(int, int, func(int, int) int) int = action

		fmt.Println("result - add", result(2, 3, addAndReturn))
		fmt.Println("result - multiply", result(2, 3, multiplyAndReturn))

		var fn func(int, int) int = selectFn(1) // fn будет выполнять то фунцию, которую вернет selectFn
		var fn2 func(int, int) int = selectFn(2)

		fmt.Println(fn(2, 3))  // 5
		fmt.Println(fn2(2, 3)) // 6
	}

	// ===========================
	// |    анонимные функции    |
	// ===========================
	fmt.Println("\n=======   анонимные функции   =======")
	{
		var f func(int, int) int = func(a, b int) int { return a + b }

		fmt.Println(f(2, 7))
		fmt.Println(f(9, 10))

		// использование анонимных функций как параметр для другой функции
		fmt.Println(action(5, 5, func(a, b int) int { return a + b })) // 10
		fmt.Println(action(5, 5, func(a, b int) int { return a * b })) // 25

	}

	// ===================
	// |    замыкание    |
	// ===================
	fmt.Println("\n=======   замыкание   =======")
	{
		fmt.Println()

		fn := outer() // n = 5, создается функция inner
		// fn теперь = функция inner, которая "ПОМНИТ" что n = 5

		fn() // 6
		fn() // 7
		fn() // 8

		fn2 := outer() // это отдельный экземпляр, его n = 5, то есть дефолтное значение

		fn2() // 6
		fn2() // 7
		fn2() // 8

		fn3 := multiply(5)

		fn3(2) // 10
		fn3(5) // 25

		fn4 := multiply(10)

		fn4(2) // 20
		fn4(5) // 50
	}

	// ===================
	// |    указатели    |
	// ===================
	fmt.Println("\n=======   указатели   =======")
	{
		fmt.Println()

		var originalString string = "original string"

		changeStr(originalString)

		fmt.Println(originalString) // originalString не изменился, потому что в методе changeStr мы изменили его копию

		changeOriginalStr(&originalString)

		fmt.Println(originalString) // originalString изменился, потому что мы указали на его оригинал (&originalString)
	}

	// ====================
	// |    псевдонимы    |
	// ====================
	fmt.Println("\n=======   псевдонимы   =======")
	{
		fmt.Println()

		var distanceMile mile = 50
		var distanceKm kilometer = 80

		distanceToObjectInMile(distanceMile)
		//distanceToObjectInMile(distanceKm) // ошибка, метод distanceToObject ожидает int в милях
		distanceToObjectInKm(distanceKm)
	}
	// ===============
	// |    срезы    |
	// ===============
	fmt.Println("\n=======   срезы   =======")
	{
		fmt.Println()

		var users []string = []string{"Kate", "Alice", "Eugen"} // рамзер не указываем

		// можно создать срез из другой последовательности элементов, например, из массива.
		// для этого применяется оператор среза s[i:j]

		var initialUsers = []string{"Bob", "Alice", users[0], "Sam", "Tom", "Paul", "Mike", "Robert"}

		// первое значение включительно второе нет
		users1 := initialUsers[:2]  // 0 и 1 то есть [Bob Alice]
		users2 := initialUsers[2:4] // 2 и 3 то есть [Kate Sam]
		users3 := initialUsers[1:7] // 1 и 6 то есть [Alice Kate Sam Tom Paul Mike]

		fmt.Println(users1)
		fmt.Println(users2)
		fmt.Println(users3)

		// также можно создать срез из строки
		var str string = "Hello World"

		slice := str[6:]

		fmt.Println(slice)

		// удаление элмемента из среза
		var array = []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

		n := 3 // индекс удаляемого элемента (Sam)

		array = append(array[:n], array[n+1:]...)
		// array[:n] - берем всё ДО значения n
		// array[n+1:]... - берем всё ПОСЛЕ значения n (n+1), и распаковываем при помощи ..., т.к. append ожидает список элементов, а не просто срез
		// троеточие в Go раскрывает срез в отдельные элементы

		fmt.Println(array) // [Bob Alice Kate Tom Paul Mike Robert]
	}

	// ===============
	// |    карты    |
	// ===============
	fmt.Println("\n=======   карты   =======")
	{
		fmt.Println()

		var empty = map[string]int{}

		fmt.Println(empty) // map[]

		var people = map[string]int{
			"Tom":   18,
			"Alice": 33,
			"Eugen": 19,
		}

		fmt.Println(people)

		// это выражение проверяет наличие ключа "Eugen" в карте people
		if val, ok := people["Eugen"]; ok { // если ключ есть, ok - true, а val - значение найденного ключа
			fmt.Printf("Карта содержит ключ Eugen, его значение: %d\n", val)
		}

		// добавление элементов
		people["Alex"] = 31

		// удаление элементов
		delete(people, "Alice")

		// редактирование элементов
		people["Tom"] = 25 // если элемент Tom отсутствует, он добавится

		fmt.Println(people)
	}
}

// если несколько параметров подряд имеют один и тот же тип, то мы можем указать тип только для последнего параметра
func add(a, b int) {
	fmt.Println()
	var c int = a + b
	fmt.Println(a, "+", b, "=", c)
}

// функции типа func(int, int) int
func addAndReturn(a int, b int) int      { return a + b }
func multiplyAndReturn(a int, b int) int { return a * b }

func addMoreValues(values ...int) {
	fmt.Println()
	var sum int = 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
		if i != len(values)-1 {
			fmt.Print(values[i], " + ")
		} else {
			fmt.Print(values[i], " = ", sum)
		}
	}
	fmt.Println()
}

func getMinValue(nums []int) int {
	fmt.Println()
	var minValue int = 9223372036854775807

	for i := 0; i < len(nums); i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
	}

	return minValue
}

// функции могут принимать в качестве параметров другие функции
func action(a int, b int, operation func(int, int) int) int { return operation(a, b) }

// функция как результат другой функции
func selectFn(index int) func(int, int) int {
	switch index {
	case 1:
		return addAndReturn
	case 2:
		return multiplyAndReturn
	}
	return addAndReturn
}

// замыкание
func outer() func() {
	var n int = 5

	inner := func() {
		n += 1
		fmt.Println(n)
	}
	return inner
}

func multiply(n int) func(int) int {
	return func(m int) int {
		println(n * m)
		return n * m
	}
}

func changeStr(str string) {
	_ = str // это нужно чтобы избежать предупреждения SA4009, компилятор будет думать что мы использовали str
	str = "changed original"
}
func changeOriginalStr(str *string) {
	*str = "changed string"
}

func distanceToObjectInMile(distanceMile mile) {
	fmt.Println("дистанция до объекта в милях:", distanceMile)
}
func distanceToObjectInKm(distanceKm kilometer) {
	fmt.Println("дистанция до объекта в километрах:", distanceKm)
}
