package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkrim(x string) int { // Функция проверяет правильность введенного римского числа
	rimtoarab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	if n, ok := rimtoarab[x]; ok {
		return n
	} else {
		return 0
	}
}

func checkint(x string) int { // Функция проверяет является ли введенное значение числом и правильность введенного числа
	n, err := strconv.Atoi(x)
	if err != nil {
		return 0
	} else if n < 1 || n > 10 {
		panic("Одно из арабских чисел не соответствует требуемому диапазону")
	} else {
		return n
	}
}

func checkznak(x string) bool { // Функция проверяет введенный знак математической операции
	sign := [4]string{"/", "+", "-", "*"}
	for _, i := range sign {
		if x == i {
			return true
		}
	}
	return false
}

func calculate(n1, n2 int, zn string) int { // Функция проводит математическую операцию
	var result int
	switch zn {
	case "/":
		result = n1 / n2
	case "*":
		result = n1 * n2
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	}
	return result
}

func arabtorim(x int) string { // Функция переводит результат вычисления функции calculate() из арабской с.с. в римскую с.с.
	arabtorimten := map[int]string{
		1:  "X",
		2:  "XX",
		3:  "XXX",
		4:  "XL",
		5:  "L",
		6:  "LX",
		7:  "LXX",
		8:  "LXXX",
		9:  "XC",
		10: "C",
	}

	arabtorimone := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	var rimnum string
	if x < 1 || x > 100 {
		panic("Невозможно вывести результат т.к в римской системе нет нуля или отрицательных чисел")
	} else {
		ten := x / 10
		one := x % 10
		if ten != 0 {
			rimnum += arabtorimten[ten]
		}
		if one != 0 {
			rimnum += arabtorimone[one]
		}
	}
	return rimnum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	vvod := strings.Split(text, " ")
	if len(vvod) < 3 {
		panic("Определенно не является математической операцией")
	} else if len(vvod) > 3 {
		panic("Возможно является математической операцией, но не удовлетворяет заданию - два операнда и один оператор (+, -, /, *)")
	}
	num1, znak, num2 := vvod[0], vvod[1], vvod[2]

	if checkznak(znak) == true {
		arab1 := checkint(num1)
		arab2 := checkint(num2)
		if arab1 == 0 && arab2 == 0 {
			rim1 := checkrim(num1)
			rim2 := checkrim(num2)
			if rim1 == 0 || rim2 == 0 {
				panic("Как минимум одно из римских чисел не является таковым или не соответствует заданному диапазону")
			} else {
				fmt.Print(arabtorim(calculate(rim1, rim2, znak)))
			}
		} else if (arab1 != 0 && checkrim(num2) != 0) || (checkrim(num1) != 0 && arab2 != 0) {
			panic("Используются одновременно разные системы счисления")
		} else {
			fmt.Print(calculate(arab1, arab2, znak))
		}
	} else {
		panic("Нераспознан введенный знак, не является возможной мат. операцией!")
	}
}
