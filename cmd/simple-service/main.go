package main

import (
	"log"
)

func fizzbuzz(input int) string {
	if (fizz(input) + buzz(input)) == "" {
		return ""
	} else {
		return fizz(input) + buzz(input)
	}
}

func fizz(input int) string {
	if input%3 == 0 {
		return "fizz"
	}
	return ""
}

func buzz(input int) string {
	if input%5 == 0 {
		return "buzz"
	}
	return ""
}

func main() {
	log.Println("beep boop")
}
