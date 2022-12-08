package main

import "fmt"

func main() {
	arr := []string{"i", "am", "stupid", "and", "weak"}
	for i, v := range arr {
		switch v {
		case "stupid":
			arr[i] = "smart"
		case "weak":
			arr[i] = "strong"
		}
	}

	for _, v := range arr {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
}
