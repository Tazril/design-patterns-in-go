package observer

import (
	"fmt"
	"testing"
)

type MockObserver struct {
}

var timesCalled = 0

func (o *MockObserver) Update(value int) {
	timesCalled += 1
	fmt.Printf("Observer Update method called with value=%v\n", value)
}

func TestCounter_Increment(t *testing.T) {
	// given
	c := Counter{}
	o := MockObserver{}
	c.Register(&o)
	timesCalled = 0
	// when
	c.Increment()
	c.Decrement()
	c.Increment()
	// then
	if timesCalled != 3 {
		t.Errorf("Got times called = %v, Expected times called = %v", timesCalled, 3)
	}
	if c.count != 1 {
		t.Errorf("Got count = %v, Expected count = %v", c.count, 1)
	}
}
