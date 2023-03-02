package template

import (
	"fmt"
	"testing"
)

type MockTeaMaker struct {
}

var addTeaCalled = false

func (rtm *MockTeaMaker) AddTea() {
	addTeaCalled = true
	fmt.Println("Adding 2-3 tea leaves")
}

var addHotWaterCalled = false

func (rtm *MockTeaMaker) AddHotWater() {
	addHotWaterCalled = true
	fmt.Println("Adding 1 cup of hot water")
}

var addMilkCalled = false

func (rtm *MockTeaMaker) AddMilk() {
	addMilkCalled = true
	fmt.Println("No milk added")
}

var addSugarCalled = false

func (rtm *MockTeaMaker) AddSugar() {
	addSugarCalled = true
	fmt.Println("Adding 1 teaspoon of sugar")
}

func (rtm *MockTeaMaker) Build() {
	rtm.AddTea()
	rtm.AddHotWater()
	rtm.AddMilk()
	rtm.AddSugar()
}

func Test_Build(t *testing.T) {
	// given
	mtm := MockTeaMaker{}
	addTeaCalled = false
	addHotWaterCalled = false
	addMilkCalled = false
	addSugarCalled = false
	// when
	mtm.Build()
	// then
	if !addTeaCalled {
		t.Errorf("AddTea() was not called")
	}
	if !addHotWaterCalled {
		t.Errorf("AddHotWater() was not called")
	}
	if !addMilkCalled {
		t.Errorf("AddMilk() was not called")
	}
	if !addSugarCalled {
		t.Errorf("AddSugar() was not called")
	}
}
