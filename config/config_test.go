package config

import (
    "testing"
)

func TestCheck(t *testing.T) {
    err := check()
    if err != nil {
        t.Fatalf("Check Failed:", err)
    }
}

func TestRead(t *testing.T) {
    read()
    if configuration == nil {
        t.Fatalf("Read Failed")
    }
}

var value_tests = []struct {
    in, out string
} {
    {"domain", "test.com"},
    {"url_len", "4"},
    {"port", "8080"},
}

func TestValue(t *testing.T) {
    for i, test := range value_tests {
        if value, _ := Value(test.in); value!=test.out {
            t.Fatalf("at index %d, expected %s, Got %s.\n", i, test.out, value)
        }
    }
}
