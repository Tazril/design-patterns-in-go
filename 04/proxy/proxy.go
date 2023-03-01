package proxy

import "fmt"

type User struct {
	name   string
	gender rune // M, F
}

type IUserService interface {
	Create(name string, gender rune) User
}

type UserService struct {
}

func (us UserService) Create(name string, gender rune) User {
	fmt.Println("Creating User with name=" + name)
	return User{
		name:   name,
		gender: gender,
	}
}

type UserServiceProxy struct {
	userService UserService
}

func (usp UserServiceProxy) Create(name string, gender rune) User {
	fmt.Println("Proxy Logging ... Received name=" + name)
	_name := name
	if gender == 'M' {
		_name = "Mr. " + name
	} else {
		_name = "Ms. " + name
	}
	result := usp.userService.Create(_name, gender)
	fmt.Println("Proxy Logging ... Created User with name=" + name)
	return result
}
