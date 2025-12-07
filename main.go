package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// fmt.Errorf()
// errors.New()

// error
type TimeError struct {
	Time time.Time
	Err  error //io.EOF
}

func (te *TimeError) Error() string {

	return fmt.Sprintf("[%v]: %v", te.Time.Format("2006/01/02 15:04:05"), te.Err)
}

func (te *TimeError) Unwrap() error {
	return te.Err
}

func NewTimeError(err error) error {
	return &TimeError{
		Time: time.Now(),
		Err:  err,
	}
}

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName) //err = os.FileNotExist = "Error file not Exist"
	if err != nil {
		return "", NewTimeError(err)
	}
	return string(data), nil
}

// func Is(err, target error) bool {}
// io.EOF
// err == io.EOF - fail

func main() {
	_, err := ReadFile("config.txt")
	if err != nil {
		// if errors.Is(err, os.ErrNotExist) {

		// 	fmt.Println("File Created.")
		// } else {
		// 	fmt.Println("такой ошибки не знаю и не обрабатываю: ", err)
		// 	os.Exit(1)
		// }
		var myErr *TimeError
		if errors.As(err, &myErr) {
			fmt.Println(myErr.Error())
			fmt.Println(myErr.Err)
			fmt.Println(myErr.Time)

		}

	}
}
