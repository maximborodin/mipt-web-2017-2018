package main

//import "fmt"
import 	"unicode"


/*Релизовать функцию, которая принимает на вход slice из чисел и возвращает slice в котором удалены все четные числа. Имя функции: RemoveEven
Пример использования:

input := []int{0, 3, 2, 5}
result := RemoveEven(input)
fmt.Println(result) // Должно напечататься [3 5]*/

func RemoveEven(slice []int) []int{
	resSlice := make([]int, 0)
	for i := 0;i < len(slice);i++ {
		if slice[i]%2 != 0 {
			resSlice = append(resSlice, slice[i])
		}
	}
	return resSlice
}


/*Написать генератор, который будет принимать число и при вызове возвращать очередную степень этого часла. Имя генератора: PowerGenerator
Пример использования:

gen := PowerGenerator(2)
fmt.Println(gen()) // Должно напечатать 2
fmt.Println(gen()) // Должно напечатать 4
fmt.Println(gen()) // Должно напечатать 8*/

func PowerGenerator(num int) (func() int) {
	res := 1
	return func() int {
		res *= num
		return res
	}
}

/*Реализуйте функцию, которая подсчитывает число различных слов в тексте (строчка).
“Словом” считается непустая последовательность букв. Слова, отличающиеся только
регистром символов, считаются одинаковыми. Для решения этой задачи вам могут
быть полезны функции из модуля unicode. Имя функции: DifferentWordsCount
Пример использования:

fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
// Должно напечатать 2*/

func DifferentWordsCount(str string) int {
	words := make(map[string]int)
	word := ""
	for _, char := range str {
		if unicode.IsLetter(char) {
			word += (string)(unicode.ToLower(char))
		} else if word != "" {
			words[word]++
			word = ""
		}
	}
	if word != "" {
		words[word]++
		word = ""
	}
	return len(words)
}