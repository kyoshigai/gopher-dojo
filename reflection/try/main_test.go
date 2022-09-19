package main

import "testing"

func toStringPointer(s string) *string {
	return &s
}

func toIntPointer(i int) *int {
	return &i
}
func TestSetString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		p    interface{}
		s    string
		want error
	}{
		{name: "normal", p: toStringPointer("hoge"), s: "aaa", want: nil},
		{name: "fail1", p: "hoge", s: "aaa", want: errCannotSetString},
		{name: "fail2", p: 100, s: "aaa", want: errCannotSetString},
		{name: "fail3", p: toIntPointer(100), s: "aaa", want: errCannotSetString},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := setString(tt.p, tt.s); got != tt.want {
				t.Errorf(
					"got: %v, want: %v, p: %v, s: %q",
					got, tt.want, tt.p, tt.s,
				)
			}
		})
	}
}
