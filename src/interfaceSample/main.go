package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}
func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (l *Lambo) Stop() {
	fmt.Println("Stop Lambo")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}
func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"C365"}
	l.Drive()
	c.Drive()
}
