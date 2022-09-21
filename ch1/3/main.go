package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for {
		input := ""
		fmt.Scan(&input)

		if input == "END" {
			os.Exit(0)
		} else {
			iInput, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("%s is not int", input)
			} else {
				fmt.Println(">", iInput)
			}
		}
	}
}
