package app

import "testing"

func TestConvertToCamel(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{name:"test_apple"}, want:"TestApple",},
		{name: "m", args: args{name:"pass_set_apple"}, want:"PassSetApple",},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToCamel(tt.args.name); got != tt.want {
				t.Errorf("ConvertToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
