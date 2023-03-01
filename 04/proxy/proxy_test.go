package proxy

import (
	"reflect"
	"testing"
)

func TestUserServiceProxy_Create(t *testing.T) {
	type args struct {
		name   string
		gender rune
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "male",
			args: args{"john doe", 'M'},
			want: User{name: "Mr. john doe", gender: 'M'},
		},
		{
			name: "female",
			args: args{"jane doe", 'F'},
			want: User{name: "Ms. jane doe", gender: 'F'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxy := UserServiceProxy{
				userService: UserService{},
			}
			if got := proxy.Create(tt.args.name, tt.args.gender); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got NewMobileOS() = %v, wanted %v", got, tt.want)
			}
		})
	}
}
