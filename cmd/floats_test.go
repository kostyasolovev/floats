package main

import "testing"

func TestComission(t *testing.T) {
	type tcase struct {
		input string
		expd  uint64
		// err   bool
	}

	var testdata = []tcase{
		{"0.01", 1},
		{"99.99", 9999},
		{"text", 0},
		{"5.5", 550},
		{"90.00", 9000},
		{"0.00", 0},
		{"4", 400},
		{"0,01", 1},
		{"9. 01", 901},
		{"9  ", 900},
		{" 80", 8000},
		{"9999", 9999},
		{" 1111", 9999},
	}

	for i, v := range testdata {
		if res := check_commission(v.input); res != v.expd {
			t.Fatalf("testcase %d: expected %d, got %d", i, v.expd, res)
		}
	}
}
