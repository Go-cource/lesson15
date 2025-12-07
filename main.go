package main

import (
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
	return fmt.Sprintf("%v: %v", te.Time.Format("2006/01/02 15:04:05"), te.Err)
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
		return "", NewTimeError(err) // TimeError {Time: 11:26, err: os.FileNotExist}
	}
	return string(data), nil
}

func main() {
	_, err := ReadFile("config.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
