package state

import "fmt"

type OrderState interface {
	Create() OrderState  // implemented by OrderCreateState
	Pack() OrderState    // implemented by OrderPackState
	Deliver() OrderState // implemented by OrderDeliverState
	State() string
}

type OrderCreateState struct {
}

func (s OrderCreateState) Create() OrderState {
	fmt.Println("Creating Order")
	return &OrderPackState{}
}

func (s OrderCreateState) Pack() OrderState {
	panic("Not Implemented")
}

func (s OrderCreateState) Deliver() OrderState {
	panic("Not Implemented")
}

func (s OrderCreateState) State() string {
	return "Order Creation State"
}

type OrderPackState struct {
}

func (s OrderPackState) Create() OrderState {
	panic("Not Implemented")
}

func (s OrderPackState) Pack() OrderState {
	fmt.Println("Packing Order")
	return &OrderDeliverState{}
}

func (s OrderPackState) Deliver() OrderState {
	panic("Not Implemented")
}

func (s OrderPackState) State() string {
	return "Order Packing State"
}

type OrderDeliverState struct {
}

func (s OrderDeliverState) Create() OrderState {
	panic("Not Implemented")

}

func (s OrderDeliverState) Pack() OrderState {
	panic("Not Implemented")
}

func (s OrderDeliverState) Deliver() OrderState {
	fmt.Println("Delivering Order")
	return nil
}

func (s OrderDeliverState) State() string {
	return "Order Delivery State"
}
