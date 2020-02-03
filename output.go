package main

import (
	"fmt"
)

type Outputer interface {
	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
}

func NewOutputer(verbose bool) Outputer {
	if verbose {
		return &standardOutputer{}
	}
	return &silentOutputer{}
}

type standardOutputer struct{}

func (s *standardOutputer) Print(a ...interface{}) (int, error) {
	return fmt.Print(a...)
}

func (s *standardOutputer) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a...)
}

func (s *standardOutputer) Println(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type silentOutputer struct{}

func (s *silentOutputer) Print(a ...interface{}) (n int, err error) {
	return 0, nil
}

func (s *silentOutputer) Printf(format string, a ...interface{}) (n int, err error) {
	return 0, nil
}

func (s *silentOutputer) Println(a ...interface{}) (n int, err error) {
	return 0, nil
}
