package organism

import "fmt"

type Human string

func (h Human) Say(words string) {
	fmt.Println(words)
}
