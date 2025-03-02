package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	PhoneBook := make(map[string]string)
	for i, key := range names {
		if PhoneBook[key] == "" {
			PhoneBook[key] = numbers[i]
		}
	}
	return PhoneBook
}

func main() {
	fmt.Println(`Program that asks a user to enter names and phone number of your contacts from phone_numbers.txt,
	stores them in a map, and prints the map.`)

	var names, numbers []string

	f, err := os.Open("phone_numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	isname := true
	for s.Scan() {
		if isname {
			names = append(names, s.Text())
			isname = false
		} else {
			numbers = append(numbers, s.Text())
			isname = true
		}
	}
	fmt.Println(createPhoneBook(names, numbers))
}
