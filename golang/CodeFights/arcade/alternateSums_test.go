package arcade

import (
	"reflect"
	"testing"
)

func Test_alternatingSums(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "[180, 105]", args: args{a: []int{50, 60, 60, 45, 70}}, want: []int{180, 105}},
		{name: "[100, 50]", args: args{a: []int{100, 50}}, want: []int{100, 50}},
		{name: "[80, 0]", args: args{a: []int{80}}, want: []int{80, 0}},
		{name: "[100, 50, 50, 100]", args: args{a: []int{100, 50, 50, 100}}, want: []int{150, 150}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingSums(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alternatingSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
