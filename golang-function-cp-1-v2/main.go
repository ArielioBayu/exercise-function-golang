package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {

	var monthMap = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	formatted_Month := monthMap[month]
	// fmt.Println(formatted_Month)

	//	Menambah angka di depan day
	formatted_Day := fmt.Sprintf("%02d", day)
	// fmt.Println(formatted_Day)

	result := fmt.Sprintf("%s-%s-%d", formatted_Day, formatted_Month, year)

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
	fmt.Println(DateFormat(1, 1, 2020))
	fmt.Println(DateFormat(31, 12, 2020))
}
