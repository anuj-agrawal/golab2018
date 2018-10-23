package main

import "testing"

func BenchmarkValidUserName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validUserName("abc")
	}
}

func TestValidUserName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{"empty", "", false},
		{"beginWithNumber", "0a", false},
		{"beginWithUnderScore", "_a", true},
		// ... more tests
		{"singleCharValid", "a", true},
		{"singleCharDollar", "$", false},
		{"endWithDollar", "a$", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if validUserName(test.in) != test.out {
				t.Fail()
			}
		})
	}
}
