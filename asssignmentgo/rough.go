package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var order string
	var quantity int
	var placed = []int{}
	counter := 0
	for 1 != 0 {
		fmt.Println("---------M  E  N  U----------")
		fmt.Println("C: coffee, 50rs ")
		fmt.Println("D: dosa, 90 rs")
		fmt.Println("T: tomato soup, 30rs")
		fmt.Println("I : idli , 30rs ")
		fmt.Println("V : vada, 35rs")
		fmt.Println("B: bhature &chhole, 40rs")
		fmt.Println("P: paneer pakoda, 40rs")
		fmt.Println("M: manchurian, 100rs")
		fmt.Println("H: hakka noodle, 80rs")
		fmt.Println("F: French fries, 40rs")
		fmt.Println("J: jalebi, 40rs")
		fmt.Println("L: lemonade, 25rs")
		fmt.Println("S: spring roll, 30rs")
		fmt.Println("-----------------------")
		fmt.Println("Press END to close counter and get total amount earned ")
		fmt.Println("Pick the order please ")
		fmt.Scan(&order)
		order = strings.ToUpper(order)

		if order == "END" {
			totalincome(placed)
		}
		fmt.Println("please enter the quantity of what u picked!!!")
		fmt.Scanln(&quantity)

		bill := bill(quantity, order)
		fmt.Println("==========================")
		fmt.Println("Total bill= ", bill)
		fmt.Println("==========================")
		placed = append(placed, bill)
		counter += 1

	}
}

func totalincome(placed []int) {
	total := 0

	for i := 0; i < len(placed); i++ {
		total = placed[i] + total
	}
	fmt.Println("Total Income for the day is : ", total)
	os.Exit(0)
}

func bill(quantity int, order string) int {
	m := map[string]int{
		"C": 50,
		"D": 90,
		"T": 30,
		"I": 30,
		"V": 35,
		"B": 40,
		"P": 40,
		"M": 100,
		"H": 80,
		"F": 40,
		"J": 40,
		"L": 25,
		"S": 30,
	}
	billreturn := quantity * m[order]
	return billreturn
}
