package main

import (
	"fmt"
	"os"
)

// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

// добавьте методы Error() и NewLabelError(label string, err error) error
// ...

func main() {
	_, err := os.ReadFile("mytest.txt") //net.Dial() err-> NewLableEror("net", err)
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(err)
	// должна выводить
	// [FILE] open mytest.txt: no such file or directory
}
