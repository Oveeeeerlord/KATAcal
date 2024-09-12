package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func sum(a, b int, op string) int {
	// Функция операций: сложения, вычитания, умножения и деления. На вход получем переменные чисел a, b type int и оператор type string.
	// Результатом работы функции возвращаем резльтута примера.
	// В зависимости от переменной op выбираем необходимое действие.

	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		if b == 0 {
			panic("На ноль делить нельзя")
		} else {
			return a / b
		}
	} else {
		return 0
	}
}

func go_to_roman(n int) string {
	// Функция конвертации арабских чисел в римские. На вход получаем переменную n type int.
	// Результатом работы функции возвращаем римское число type string

	res := ""

	elements := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	// Блок сортировки карты от большого числа к меньшему
	keys := make([]int, len(elements))

	i := 0

	for k := range elements {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	slices.Reverse(keys)

	// Блок логики получения римского числа:
	// 1. Перебираем по ключам карту 2. Делим заданное число на элемент ключа;
	// 3. Добавляем в результат римскую цифру то число раз, сколько получилось в пукте 2; 3. Получаем остаток и перезаписываем переменную n;
	// 4. Повторяем операцию пока не получим римское число.
	for _, p := range keys {
		y := n / p
		res += strings.Repeat(elements[p], y)
		n %= p

	}
	return res
}

func go_to_arabic(n string) int {
	// Функция конвертации римских чисел в арабские. На вход получаем переменную n type string.
	// Результатом работы функции возвращаем арабское число type int

	elements := map[string]int{
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	// Блок логики получения арабского числа:
	// 1.Присваиваем перменной rome_num значение n; 2. Перебираем число получая индекс и римску цифру число
	// 3.Перебираем список карты получения ключ-значение; 4. Если римская буква равна ключу карты то проверяем условия;
	// 5. Если индекс равен длине строки римского чила, то суммируем в полученное число с перменной res, если нет, то выполняем ряд условий для
	// отображения корректного числа (Проверка условий, например, если рассматриваемая число равно С и последующее число равно М, то мы отнимаем значение ключа).
	rome_num := n

	res := 0

	for n, i := range rome_num {
		for ro, ar := range elements {
			if string(i) == ro {
				if n == len(rome_num)-1 {
					res += ar
				} else {
					if string(i) == "C" && string(rome_num[n+1]) == "M" {
						res -= ar
						continue
					}
					if string(i) == "C" && string(rome_num[n+1]) == "D" {
						res -= ar
						continue
					}
					if string(i) == "X" && string(rome_num[n+1]) == "C" {
						res -= ar
						continue
					}
					if string(i) == "X" && string(rome_num[n+1]) == "L" {
						res -= ar
						continue
					}
					if string(i) == "I" && string(rome_num[n+1]) == "X" {
						res -= ar
						continue
					}
					if string(i) == "I" && string(rome_num[n+1]) == "V" {
						res -= ar
						continue
					} else {
						res += ar
					}
				}
			}
		}
	}
	return res
}

func check_type(i string) (string, int) {
	// Проверка аргументов на заданные числа. Числа должны быть арабскими (от 1 до 10) или римскими (от I до X)
	// На вход получаем получаем строку с каким-то значением. Результатом работы функции является строка с арабским или римским числом и также int со значнием 0 или 1 (значение введено чтобы сравнить в дальнейшем два числа примера)

	ar_list := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rm_list := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	// Создаем два массива с арабскими и римскими цифрами
	// Проверям, чтобы тот или иной масси содержал необходимо арабское или римское число
	// Если число не содержится, то выполняется ошибка (паника), что введено неверное число.
	if slices.Contains(ar_list, i) {
		return i, 0
	} else if slices.Contains(rm_list, i) {
		return i, 1
	} else {
		panic("Введеное неверное число. Числа должны быть арабскими (от 1 до 10) или римскими (от I до X)")
	}
}

func oper_check(op string) string {
	// Функция проверки операнда
	// На вход получаем какое-то значние type string. Проверяем с помощью условных выражений, чтобы было данное значение.
	// Если значение не найдено, то выполеняется ошибка (паника), что введен неверный оператор.
	// Результатом работы функции является возвращение необходимо опертора

	if op == "+" {
		return op
	} else if op == "-" {
		return op
	} else if op == "*" {
		return op
	} else if op == "/" {
		return op
	} else {
		panic("Неверный операнд")
	}
}

func main() {

	//Блок Ввод данных пользователя в консоли
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("КАЛЬКУЛЯТОР сложения, деления, умножения, деления римских и арабских цифр")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Умеет работать с арабскими (от 1 до 10) или римскими (от I до X)")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Print("Введите пример (в формате 10 * 5): ")
	text, _ := reader.ReadString('\n')

	strp_st := strings.Fields(text) // Разбиваем строку на массив
	if len(strp_st) != 3 {          // Проверям условие, чтобы элементов массива было 3, Если значение != 3, выводим сообщение об ошибке и завершаем программу
		fmt.Println("Неверно введен пример")
	} else { // Если условие подходит, выполняем программу дальше

		first_value_str := strp_st[0] // Присваивем трем перменным значения из массива, две переменные числа и одна переменная - оператор
		second_value_str := strp_st[2]
		operand := strp_st[1]

		a_d, ir := check_type(first_value_str) // Проверяем числа
		b_d, ir2 := check_type(second_value_str)

		o_c := oper_check(operand) // Проверяем оператор

		if ir == 0 && ir2 == 0 { // Выполняем условие, если два числа равны арабским числаv, то сразу выполняем операцию
			to_digit, err := strconv.Atoi(a_d) // Переводим арабское число из string в int
			if err != nil {
				panic(err)
			}
			to_digit2, err := strconv.Atoi(b_d)
			if err != nil {
				panic(err)
			}

			itog := sum(to_digit, to_digit2, o_c)
			fmt.Println(itog) // Выводим получивший результат в консоль

		} else if ir == 1 && ir2 == 1 { // Выполняем условие, если два числа равны римским числам, то выполняем ряд действий:
			s1 := go_to_arabic(a_d) // 1. Конвертируем число в арабское
			s2 := go_to_arabic(b_d)
			itog_ar := sum(s1, s2, o_c) // 2. Выполняем операцию
			if itog_ar < 0 {            // 3. проверям, чтобы римское число не было отрицательным
				panic("Ошибка: Итог уравнения римских цифр является отрциательное число")
			} else {
				itog := go_to_roman(itog_ar) // 4. Переводим число обратно в римское
				fmt.Println(itog)            // 5. Выводим результат в консоль
			}
		} else {
			panic("Неверное значение")
		}

	}

}