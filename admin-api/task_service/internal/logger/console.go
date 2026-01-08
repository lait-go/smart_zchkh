package logger

import "fmt"

type Log interface {
	Log(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Console struct{}

func (c *Console) Log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("[INFO]", msg)
}

func (c *Console) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("[ERROR]", msg)
}