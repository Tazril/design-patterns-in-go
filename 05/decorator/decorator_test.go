package decorator

import (
	"testing"
)

func TestMargheritaPizza(t *testing.T) {
	tp := ThickCrustPizza{}     // price = 150
	cp := CheeseAddonPizza{tp}  // price = 80
	op := OreganoAddonPizza{cp} // price = 40
	if op.Price() != 270 {
		t.Errorf("Got Price %v Expected Price %v", op.Price(), 270)
	}
}

func TestDesiPizza(t *testing.T) {
	tp := ThinCrustPizza{}       // price = 100
	ep := EggplantAddonPizza{tp} // price = 50
	top := TomatoAddonPizza{ep}  // price = 20
	if top.Price() != 170 {
		t.Errorf("Got Price %v Expected Price %v", top.Price(), 170)
	}
}
