package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func romToDecimal(rom string) int {
	romKeys := map[byte]int{'I': 1, 'V': 5, 'X': 10}
	res := 0
	last := 0
	for i := len(rom) - 1; i >= 0; i-- {
		value := romKeys[rom[i]]
		if value < last {
			res -= value
		} else {
			res += value
		}
		last = value
	}
	return res
}

func IsRoman(s string) bool {
	for _, char := range s {
		if !strings.ContainsRune("IVX", char) {
			return false
		}
	}
	return true
}

func _Input(input string) (int, bool) {
	if IsRoman(input) {
		number := romToDecimal(input)
		if number < 1 || number > 10 {
			panic("Числа вне допустимого диапаззона ")
		}
		return number, true
	} else {
		number, err := strconv.Atoi(input)
		if err != nil || number < 1 || number > 10 {
			panic("числа вне допустимого диапазона")
		}
		return number, false
	}
}

func decimalToRoman(num int) string {
	if num < 1 {
		panic("Римские числа не могут быть < I")
	}
	values := []int{10, 9, 5, 4, 1}
	keys := []string{"X", "IX", "V", "IV", "I"}
	res := ""
	for i, value := range values {
		for num >= value {
			num -= value
			res += keys[i]
		}
	}
	return res
}
func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Неверный оператор")
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}
	num1, isRoman1 := _Input(parts[0])
	num2, isRoman2 := _Input(parts[2])
	if isRoman1 != isRoman2 {
		panic("Числа должны быть римскими или арабскими")
	}
	res := calculate(num1, num2, parts[1])
	if isRoman1 {
		fmt.Printf("Результат: %s\n", decimalToRoman(res))
	} else {
		fmt.Printf("Результат: %d\n", res)
	}
}
