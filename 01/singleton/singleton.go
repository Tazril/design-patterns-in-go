package singleton

type Singleton struct {
	// count example field
	count int
}

// AddOne example method
func (singleton *Singleton) AddOne() {
	singleton.count += 1
}

var instance *Singleton

// GetSingletonInstance Returns the same instance when called
func GetSingletonInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
