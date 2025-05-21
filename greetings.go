package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Holam(nombre string) (string, error) {
	if nombre == "" {
		return "", errors.New("nombre vacio")
	}
	message := fmt.Sprintf(Randomformat(), nombre)
	return message, nil
}

func Holas(nombres []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, nombre := range nombres {
		message, err := Holam(nombre)
		if err != nil {
			return nil, err
		}

		messages[nombre] = message
	}

	return messages, nil
}

func Randomformat() string {
	formats := []string{
		"Hola %v bienvenido",
		"Que bueno verte %v",
		"Como estas %v encargado de conocerte",
	}
	return formats[rand.Intn(len(formats))]
}
