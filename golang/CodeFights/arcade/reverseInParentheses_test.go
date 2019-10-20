package arcade

import "testing"

func Test_reverseInParentheses(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "(bar)", args: args{inputString: "(bar)"}, want: "rab"},
		{name: "foo(bar)baz", args: args{inputString: "foo(bar)baz"}, want: "foorabbaz"},
		{name: "foo(bar)baz(blim)", args: args{inputString: "foo(bar)baz(blim)"}, want: "foorabbazmilb"},
		{name: "foo(bar(baz))blim", args: args{inputString: "foo(bar(baz))blim"}, want: "foobazrabblim"},
		{name: "()", args: args{inputString: "()"}, want: ""},
		{name: "(abc)d(efg)", args: args{inputString: "(abc)d(efg)"}, want: "cbadgfe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInParentheses(tt.args.inputString); got != tt.want {
				t.Errorf("reverseInParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
