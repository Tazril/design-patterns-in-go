# Factory Design Pattern
The Factory design pattern is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. This pattern is useful when you want to decouple object creation from object usage.

## Code Highlight
The following Go code implements the Factory design pattern:
```
func NewMobileOS(tag string) IMobileOS {
	switch tag {
	case "opensource":
		return &Android{}
	case "elite":
		return &IOS{}
	default:
		return nil
	}
}
```

## Usage
The Factory pattern can be used in situations where you want to decouple object creation from object usage. One example of where this pattern might be used is in a mobile operating system. You may want to create different types of mobile operating systems (such as Android and iOS), but you don't want to hard-code the creation of these objects in your program.

By using the Factory pattern, you can create an interface (IMobileOS) for creating mobile operating systems, and then create subclasses (Android and IOS) that implement this interface. You can then use the Factory method (NewMobileOS) to create objects based on a tag, which hides the implementation of object creation from the rest of the program.

Another example of where the Factory pattern might be used is in a web application. You may want to create different types of web pages (such as home pages and product pages), but you don't want to hard-code the creation of these objects in your program.

Overall, the Factory pattern is a useful pattern for situations where you want to decouple object creation from object usage. It can help to simplify the design of your program and make it easier to maintain.
