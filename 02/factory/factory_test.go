package factory

import (
	"reflect"
	"testing"
)

func TestNewMobileOS(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IMobileOS
	}{
		{
			name: "android",
			args: args{t: "opensource"},
			want: &Android{},
		},
		{
			name: "ios",
			args: args{t: "elite"},
			want: &IOS{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMobileOS(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got NewMobileOS() = %v, wanted %v", got, tt.want)
			}
		})
	}
}
