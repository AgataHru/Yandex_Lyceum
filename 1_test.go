package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		input     []byte
		exp       bool
		exp_count int
	}{
		{
			input:     []byte("aboba"),
			exp:       false,
			exp_count: 5,
		},
		{input: []byte{0xff}, exp: true, exp_count: 0},
	}
	for i, tc := range cases {
		got_count, err := GetUTFLength(tc.input)
		check := false
		if err != nil {
			check = true
		}
		if got_count != tc.exp_count || check != tc.exp {
			t.Fatalf("#%d, expected count: %d, but got: %d, error: %s ", i, tc.exp_count, got_count, err)
		}
	}
}
