package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []structs{
		name		string
		inputURL	string
		expected	string
	}{
		{
			name: "one"
		},
		{
			name: "two"
		}
	}

	for i, tc := range tests {
		t.Run(tc.name func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			
		})
	}
}

func TestGetStocksFromHTML(t *testing.T) {
	tests := []structs{
		name 			string
		inputURL		string
		inputBody		string
		expected		[]string
		errorContains	string
	}{
		{
			name: "one"
		},
		{
			name: "two"
		}
	}
	
	for i, tc := range tests {
		t.Run(tc.name func(t *testing.T) {
			actual, err := getStocksFromHTML(tc.inputBody, tc.inputURL)
			
		})
	}
}
