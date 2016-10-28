package manager

import (
	"testing"
	"fmt"
)

var testData = []struct {
	in string
	out string
} {
	{"http://", ""},
	{"https://", ""},
	{"", ""},
	
	{"ftp://", ""},
	{"ftps://", ""},
	
	{"http://google", ""},
	{"https://google", ""},
	{"ftp://google", ""},
	{"ftps://google", ""},
	
	{"http://google.com", "http://google.com"},
	{"https://google.com", "https://google.com"},
	{"ftp://google.com", "ftp://google.com"},
	{"ftps://google.com", "ftps://google.com"},
	
	{"google.com", "http://google.com"},
	{"google", ""},
}

func TestFormatUrl(t *testing.T) {
	for _, tc := range testData {
		url, _ := FormatUrl(tc.in)
		
		if tc.out != url {
			t.Error("Test Failed. Input:", tc.in, "Result:", url, "Expected:", tc.out) 
		} else {
			fmt.Println("Test Passed.", "Input:", tc.in, "Output:", url)
		}		
	}
}
