package builder

import (
	"reflect"
	"testing"
)

func TestPhoneBuilder_Build(t *testing.T) {
	type args struct {
		t map[string]any
	}
	tests := []struct {
		name string
		args args
		want *Phone
	}{
		{
			name: "model_name_ram",
			args: args{map[string]any{"SetModel": "redmi", "SetName": "X200", "SetRAM": int8(16)}},
			want: &Phone{model: "redmi", name: "X200", ramInGB: 16, os: "Android", screenSize: 4.5},
		},
		{
			name: "complete",
			args: args{map[string]any{"SetModel": "redmi", "SetName": "X200", "SetScreenSize": float32(6.5), "SetRAM": int8(8), "SetOS": "Android"}},
			want: &Phone{model: "redmi", name: "X200", ramInGB: 8, os: "Android", screenSize: 6.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := &PhoneBuilder{}
			for k, v := range tt.args.t {
				reflect.ValueOf(builder).MethodByName(k).Call([]reflect.Value{reflect.ValueOf(v)})
			}
			if got := builder.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got NewMobileOS() = %v, wanted %v", got, tt.want)
			}
		})
	}
}
