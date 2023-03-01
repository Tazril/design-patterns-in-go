package decorator

type IPizza interface {
	Price() int
}

type ThinCrustPizza struct {
	pizza IPizza
}

type ThickCrustPizza struct {
	pizza IPizza
}

type CheeseAddonPizza struct {
	pizza IPizza
}

type OreganoAddonPizza struct {
	pizza IPizza
}

type TomatoAddonPizza struct {
	pizza IPizza
}

type EggplantAddonPizza struct {
	pizza IPizza
}

func pizzaPrice(pizza IPizza) int {
	if pizza != nil {
		return pizza.Price()
	}
	return 0
}

func (p ThinCrustPizza) Price() int {
	return pizzaPrice(p.pizza) + 100
}
func (p ThickCrustPizza) Price() int {
	return pizzaPrice(p.pizza) + 150
}
func (p CheeseAddonPizza) Price() int {
	return pizzaPrice(p.pizza) + 80
}
func (p OreganoAddonPizza) Price() int {
	return pizzaPrice(p.pizza) + 40
}
func (p TomatoAddonPizza) Price() int {
	return pizzaPrice(p.pizza) + 20
}
func (p EggplantAddonPizza) Price() int {
	return pizzaPrice(p.pizza) + 50
}
