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
	Err  error
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

func main() {
	_, err := ReadFile("config.txt")
	if err != nil {
		fmt.Println(err) //с временем
		fmt.Printf("Изначальная ошибка: %v\n", errors.Unwrap(err))
		os.Exit(1)
	}
}
