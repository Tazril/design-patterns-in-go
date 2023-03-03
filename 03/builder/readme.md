# Builder Design Pattern
The Builder design pattern is a creational design pattern that separates the construction of a complex object from its representation so that the same construction process can create different representations.

## Code Highlight
The following Go code implements the Builder design pattern:

```
type Phone struct {
	model      string
	name       string
	os         string
	ramInGB    int8
	screenSize float32
}

type PhoneBuilder struct {
	phone Phone
}

// Create Setters
func (b *PhoneBuilder) SetModel(model string) *PhoneBuilder {
	b.phone.model = model
	return b
}

func (b *PhoneBuilder) SetName(name string) *PhoneBuilder {
	b.phone.name = name
	return b
}

// ... more

func (b *PhoneBuilder) Build() *Phone {
        // Pass defaults
	if b.phone.model == "" {
		b.phone.model = "Unknown"
	}
	if b.phone.name == "" {
		b.phone.name = "Unknown"
	}
	
        // ... more
  
	return &b.phone
}
```

## Usage
The Builder pattern can be used when you want to create a complex object step by step, by calling a series of methods on a builder object. Each method sets one or more properties of the object being built.

In the above code snippet, the PhoneBuilder struct has a number of methods (SetModel, SetName, SetOS, SetRAM, SetScreenSize) that can be called to set the properties of a Phone object. The Build method constructs the Phone object and returns it.

Using the Builder pattern can make it easier to create complex objects by breaking down the construction process into smaller steps. It also allows you to create different representations of the same object, by calling different combinations of the builder methods.

For example, you might use the Builder pattern to create different types of phones with different properties. You could call the SetModel, SetName, SetOS, SetRAM, and SetScreenSize methods to set the properties of the phone, and then call the Build method to create the phone object.

Overall, the Builder pattern is a useful pattern for situations where you want to create a complex object step by step, and where you may want to create different representations of the same object. It can help to simplify the construction of complex objects and make it easier to maintain your code.






