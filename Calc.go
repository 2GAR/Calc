//////////////////////////////////////////////////////
// Ввод Данных через пробел (3 + 5, VI - III и т.д.)//
//////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, y, ope, mode string
	for {
		digits := strings.Fields(Scan1())
		switch {
		case len(digits) != 3:
			fmt.Println("Ошибка ввода! Неверный формат математической операции")
			os.Exit(1)
		case len(digits) == 3:
			x = digits[0]
			ope = digits[1]
			y = digits[2]

		}
		mode = checkSystem(x, y)
		if mode == "rome" {
			var a, b int
			a, b = fromRometoInt(x, y)
			res := maths(a, b, ope)
			if res < 0 {
				fmt.Println("В римской системе нет отрицательных чисел")
				os.Exit(1)
			}
			fmt.Println("Результат")
			fmt.Println(toRome(res))
		} else if mode == "arab" {
			x, _ := strconv.Atoi(x)
			y, _ := strconv.Atoi(y)
			fmt.Println("Результат:")
			fmt.Println(maths(x, y, ope))
		} else if mode == "none" && mode != "usederror" {
			fmt.Println("Ошибка, Допустим ввод чисел от 1 до 10")
			os.Exit(1)
		}
	}
}

// Функция перевода из Римской в Арабскую
func fromRometoInt(a, b string) (int, int) {
	var x, y int
	rome := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i <= 9; i++ {
		if a == rome[i] {
			x = i + 1
		}
		if b == rome[i] {
			y = i + 1
		}
	}
	return x, y
}

// Фунция считывания строки
func Scan1() string {
	fmt.Println("Введите математическое выражение:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

// Функция перевода в Римскую систему счислений
func toRome(res int) string {
	roman := ""
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, conversion := range conversions {
		for res >= conversion.value {
			roman += conversion.digit
			res -= conversion.value
		}
	}
	return roman
}

// Функция сверка Римских и Арабских цифр
func checkSystem(x, y string) string {
	modex := "none"
	modey := "none"
	mode := "none"
	rome := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arab := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i := 0; i <= 9; i++ {
		if x == rome[i] {
			modex = "rome"
		}
		if x == arab[i] {
			modex = "arab"
		}
		if y == rome[i] {
			modey = "rome"
		}
		if y == arab[i] {
			modey = "arab"
		}
	}
	if (modex == modey) && (modex != "none") {
		mode = modex
	} else if (modex != modey) && (modex != "none") && (modey != "none") {
		fmt.Println("Ошибка использования разных систем счисления")
		mode = "usederror"
		os.Exit(1)
	}
	return mode
}

// Функция Калькулятор
func maths(x, y int, ope string) int {
	if ope == "+" {
		return x + y
	} else if ope == "-" {
		return x - y
	} else if ope == "x" || ope == "*" {
		return x * y
	} else if ope == "/" {
		return x / y
	} else if ope == "^" || ope == "**" {
		return x ^ y
	}
	return 1
}
