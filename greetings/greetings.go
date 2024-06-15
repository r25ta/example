package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//IF NO NAME WAS GIVEN, RETURN AN ERROR WITH A MESSAGE
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hello return a map tha associates each of the named people
// With a greeting message
func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages.
	messages := make(map[string]string)

	//Loop through the received slice of names, calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	//In the map, associate the retrieved message with the name.
	return messages, nil
}

// randomFormat returns one of a set of greeting message.
// The returned message is selected at random.
func randomFormat() string {
	//A slice of message format and initialized with values
	formats := []string{
		"Hi, %v Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
		"Whats up, %v ?",
	}

	//fmt.Println(len(formats))
	randomIndex := rand.Intn(len(formats))
	//fmt.Println(randomIndex)
	//Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[randomIndex]
}
