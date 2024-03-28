package main

import (
	"fmt"
)

func namaTerpendek(name1 string, name2 string) string {
	if len(name1) < len(name2) {
		return name1
	} else if len(name1) > len(name2) {
		return name2
	} else {
		if name1 < name2 {
			return name1
		} else {
			return name2
		}
	}
}

func FindShortestName(names string) string {
	currentName := ""
	isFirstTime := true
	minName := ""

	for _, value := range names {
		if string(value) == ";" || string(value) == " " || string(value) == "," {

			if isFirstTime == true {
				minName = currentName
				isFirstTime = false
			} else {
				minName = namaTerpendek(currentName, minName)
			}
			currentName = ""
		} else {
			currentName += string(value)
		}
	}
	if currentName != "" {
		if isFirstTime == true {
			minName = currentName
			isFirstTime = false
		} else {
			minName = namaTerpendek(currentName, minName)
		}
	}
	return minName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
