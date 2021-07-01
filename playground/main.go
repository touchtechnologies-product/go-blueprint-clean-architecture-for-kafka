package main

import (
	"fmt"
)

type speaker interface {
	say() string
	bark() string
	roar() string
}

type human struct{}

func (h *human) say() string {
	return "Say something!"
}

func (h *human) bark() string {
	return "Bark!!"
}

func (h *human) roar() string {
	return "roar!!!"
}

func newSpeaker() speaker {
	return &human{}
}

func main() {
	fmt.Println("Hello, playground")
	var telephone speaker
	telephone.bark()

}
