package main

import (
	"testing"
)

func Test_functionOne(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Param",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			functionOne()
		})
	}
}

func Test_doSomething(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doSomething()
		})
	}
}

func HelloWorld() (string, error) {
	return "hi", nil
}
