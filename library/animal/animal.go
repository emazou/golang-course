package animal

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Println(d.Name + " Woof! Woof!")
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() {
	fmt.Println(c.Name + " Meow! Meow!")
}

// function to print the animal sound

func PrintSoundAnimal(a Animal) {
	a.Speak()
}
