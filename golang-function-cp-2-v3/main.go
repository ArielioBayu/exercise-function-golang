package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {

	totalVokal := 0
	totalKonsonan := 0

	for _, v := range str {
		result := strings.ToLower(string(v))
		if result == "a" || result == "i" || result == "u" || result == "e" || result == "o" {
			totalVokal += 1
		} else if result != " " && result != "," {
			totalKonsonan += 1
		}
	}
	// fmt.Println(totalVokal, totalKonsonan)
	isValid := false
	if totalVokal == 0 || totalKonsonan == 0 {
		isValid = true
	} else {
		isValid = false
	}

	return totalVokal, totalKonsonan, isValid // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("kopi"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
