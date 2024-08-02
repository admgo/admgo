package main

import (
	"errors"
)

func main() {
	err := errors.New("daw")
	err.Error()
	println(err.Error())
}
