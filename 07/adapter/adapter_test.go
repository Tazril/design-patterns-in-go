package adapter

import (
	"fmt"
	"testing"
)

var called = false

type MockPencil struct {
}

func (pa *MockPencil) Draw(text string) {
	called = true
	fmt.Println("Drawing text=" + text)
}

func TestNoOp_Sent(t *testing.T) {
	// given
	pencil := MockPencil{}
	adapter := PenAdapter{pencil: &pencil}
	writer := Writer{pen: &adapter}
	called = false
	// when
	writer.WriteText("Once upon a time")
	// then
	if !called {
		t.Errorf("Expected MockPencil Draw function to be called")
	}
}
