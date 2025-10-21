package main

import "fmt"

func main() {
	var a int
	prov := 1
	fmt.Println("Введите число")
	fmt.Scan(&a)

	for a < 12307 && prov == 1 {
		if a < 0 {
			a = a * -1
		} else if a%7 == 0 {
			a = a * 39
		} else if a%9 == 0 {
			a = a*13 + 1
		} else {
			a = (a + 2) * 3
		}

		if a%9 == 0 && a%13 == 0 {
			prov = 0
		} else {
			a++
		}
	}

	if prov == 0 {
		fmt.Println("service error")
	} else {
		fmt.Printf("итоговое число: %d", a)
	}
}
