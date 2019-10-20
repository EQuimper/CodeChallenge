package arcade

import (
	"reflect"
	"testing"
)

func Test_addBorder(t *testing.T) {
	type args struct {
		picture []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "['abc','ded']", args: args{picture: []string{"abc", "ded"}}, want: []string{"*****", "*abc*", "*ded*", "*****"}},
		{name: "['wzy**']", args: args{picture: []string{"wzy**"}}, want: []string{"*******", "*wzy***", "*******"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBorder(tt.args.picture); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
