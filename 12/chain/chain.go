package chain

import "fmt"

type Logger interface {
	Log(tag string, message string) string
}

type ErrorLogger struct {
	next Logger
}

func (l *ErrorLogger) Log(tag string, message string) string {
	if tag == "error" {
		return fmt.Sprintf("Error: %v\n", message)
	} else if l.next != nil {
		return l.next.Log(tag, message)
	}
	return ""
}

type InfoLogger struct {
	next Logger
}

func (l *InfoLogger) Log(tag string, message string) string {
	if tag == "info" {
		return fmt.Sprintf("Info: %v\n", message)
	} else if l.next != nil {
		return l.next.Log(tag, message)
	}
	return ""
}

type WarnLogger struct {
	next Logger
}

func (l *WarnLogger) Log(tag string, message string) string {
	if tag == "warn" {
		return fmt.Sprintf("Warn: %v\n", message)
	} else if l.next != nil {
		return l.next.Log(tag, message)
	}
	return ""
}
