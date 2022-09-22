package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Panic("need one or more arguments")
	}

	sum := 0.
	count := 0.
	for i := 1; i < len(args); i++ {
		float, err := strconv.ParseFloat(args[i], 1218)
		if err == nil {
			sum += float
			count++
		}
	}
	fmt.Println(sum / count)
}
