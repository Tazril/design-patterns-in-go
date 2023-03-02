package observer

type IObserver interface {
	Update(value int)
}

type IObservable interface {
	Register(observer IObserver)
	Remove()
	Notify(value int)
}

type Counter struct {
	count    int
	observer IObserver
}

func (c *Counter) Increment() {
	c.count += 1
	c.Notify(c.count)
}

func (c *Counter) Decrement() {
	c.count -= 1
	c.Notify(c.count)
}

func (c *Counter) Register(observer IObserver) {
	c.observer = observer
}

func (c *Counter) Remove() {
	c.observer = nil
}

func (c *Counter) Notify(value int) {
	if c.observer != nil {
		c.observer.Update(value)
	}
}
