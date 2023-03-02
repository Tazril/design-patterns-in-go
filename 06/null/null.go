package null

import "fmt"

type INotification interface {
	Sent(message string) bool
}

type Email struct {
}

type Sms struct {
}

// NoOp Instead of declaring INotification as nil, we will init as this when we don't want to send message
type NoOp struct {
}

func (n *Email) Sent(message string) bool {
	fmt.Println("Sending notification via email: " + message)
	return true
}

func (n *Sms) Sent(message string) bool {
	fmt.Println("Sending notification via sms: " + message)
	return true
}

func (n *NoOp) Sent(message string) bool {
	fmt.Println("Doing nothing with message: " + message)
	return false
}
