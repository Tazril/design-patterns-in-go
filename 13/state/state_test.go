package state

import (
	"testing"
)

func TestOrderLifecycle(t *testing.T) {
	// given
	var state OrderState
	state = OrderCreateState{}
	if state.State() != "Order Creation State" {
		t.Errorf("Got state=%v Expected Order Creation State at start", state.State())
		return
	}
	state = state.Create()
	if state.State() != "Order Packing State" {
		t.Errorf("Got state=%v Expected Order Packing State after Order Creation State", state.State())
		return
	}
	state = state.Pack()
	if state.State() != "Order Delivery State" {
		t.Errorf("Got state=%v Expected Order Delivery State after Order Packing State", state.State())
		return
	}
	state = state.Deliver()
	if state != nil {
		t.Errorf("Got state= %v Expected state to be nil after Order Delivery State", state.State())
		return
	}
}

func TestOrderCreateState_Deliver(t *testing.T) {
	// given
	var state OrderState
	state = OrderCreateState{}
	if state.State() != "Order Creation State" {
		t.Errorf("Got state=%v Expected Order Creation State at start", state.State())
		return
	}
	// panic test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Calling Deliver() method at Order Creation State did not cause panic")
		}
	}()
	state.Deliver()
}
