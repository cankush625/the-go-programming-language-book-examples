package main

import "fmt"

type Bird interface {
	GetCanFly() bool
	GetSound() string
	GetName() string
}

type BirdS struct {
	Name   string
	Sound  string
	CanFly bool
}

func (b BirdS) GetCanFly() bool {
	return b.CanFly
}

func (b BirdS) GetSound() string {
	return b.Sound
}

func (b BirdS) GetName() string {
	return b.Name
}

func GetBirdDetails(bird Bird) string {
	canFly := "cannot fly"
	if bird.GetCanFly() {
		canFly = "can fly"
	}
	return fmt.Sprintf("Sound of %s is %s and this bird %s.", bird.GetName(), bird.GetSound(), canFly)
}

func main() {
	peacock := BirdS{
		Name:   "Peacock",
		CanFly: true,
		Sound:  "may-awe",
	}
	// peacock is a valid Bird because it satisfies the Bird interface
	fmt.Println(GetBirdDetails(peacock)) // "Sound of Peacock is may-awe and this bird can fly."
}

// Note:
// To satisfy an interface, the type must have all the
// interface methods implemented

// The type interface{} can be used to indicate Any type.
// The function param with type interface{} can accept
// any type of value string, bool, int, struct, interface, etc.
