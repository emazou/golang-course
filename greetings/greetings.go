package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty name")
	}
	message := fmt.Sprint(randomGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}

func randomGreeting() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello, %v. Welcome!",
		"Salutations, %v. Welcome!",
	}
	return formats[rand.Intn(len(formats))]
}
