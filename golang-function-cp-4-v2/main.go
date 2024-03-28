package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	firstTime := true

	for _, value := range data {
		resultInput := strings.Contains(value, input)

		if resultInput {
			if firstTime {
				result += value
				firstTime = false
			} else {
				result += "," + value
			}
		}
	}

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
	fmt.Println(FindSimilarData("mobil", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
	fmt.Println(FindSimilarData("motor", "mobil APV", "mobil Avanza", "motor matic", "motor gede", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel"))
}
