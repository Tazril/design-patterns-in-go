package template

import "fmt"

type TeaMaker interface {
	AddTea()
	AddHotWater()
	AddMilk()
	AddSugar()
}

type RedTeaMaker struct {
}

func (rtm *RedTeaMaker) AddTea() {
	fmt.Println("Adding 2-3 tea leaves")
}

func (rtm *RedTeaMaker) AddHotWater() {
	fmt.Println("Adding 1 cup of hot water")
}

func (rtm *RedTeaMaker) AddMilk() {
	fmt.Println("No milk added")
}

func (rtm *RedTeaMaker) AddSugar() {
	fmt.Println("Adding 1 teaspoon of sugar")
}

func (rtm *RedTeaMaker) Build() {
	rtm.AddTea()
	rtm.AddHotWater()
	rtm.AddMilk()
	rtm.AddSugar()
}
