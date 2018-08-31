package main

import (
	"giggle/organism"
)

func main() {
	var lee organism.Human
	var eagle organism.Animal = "eagle"
	lee.Say("hello world.")
	eagle.Fly()
}
