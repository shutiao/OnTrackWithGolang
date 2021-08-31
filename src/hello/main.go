package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func getGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err := errors.New("Too early for greetings!")
		return message, err
	}
	if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}

func declareArray() {
	var langs [3]string
	langs[0] = "Go"
	langs[1] = "Ruby"
	langs[2] = "JavaScript"
	fmt.Println(langs)
}

func getLangs() []string {
	langs := []string{"Go", "Ruby", "JavaScript"}
	langs = append(langs, "LolCode")
	return langs
}


func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("Hello from Gopher")
	}
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < hourOfDay; i++ {
		fmt.Println("Ding Ding Ding")
	}

	j := 0
	isLessThanFive := true
	for isLessThanFive {
		fmt.Println(j)
		j++
		if j >= 5 {
			isLessThanFive = false
		}
	}
	fmt.Println(greeting)
	declareArray()

	langs := getLangs()
	for i:= range langs {
		fmt.Println(langs[i])
	}
	for _, element := range langs {
		fmt.Println(element)
	}
}
