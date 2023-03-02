package null

import (
	"testing"
)

func TestNoOp_Sent(t *testing.T) {
	var notification INotification = &NoOp{}
	sent := notification.Sent("Hi, How are you ?")
	if sent == true {
		t.Errorf("Got sent=true expected sent=false")
	}
}
