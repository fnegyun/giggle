package organism

import "fmt"

type Human string

type Animal string

type Bird interface {
	Fly()
}

func (h Human) Say(words string) {
	fmt.Println(words)
}

func (a Animal) Fly() {
	fmt.Println(a + " can fly.")
}
