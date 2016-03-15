package config

import (
    "testing"
)

func TestCheck(t *testing.T) {
    _, err := Check()
    if err != nil {
        t.Fatalf("Check Failed:", err)
    }
}

var value_tests = []struct {
    in, out string
} {
    {"domain", "test.com"},
    {"url_len", "4"},
    {"port", "8080"},
}

func TestRead(t *testing.T) {
    conf := Read()
    for _, test := range value_tests {
        if v, _ := conf[test.in]; v != test.out {
            t.Fatalf("Test failed, expected %s, got %s.\n", test.out, v)
        }
    }
}

// func TestValue(t *testing.T) {
//     for i, test := range value_tests {
//         if value, _ := Value(test.in); value!=test.out {
//             t.Fatalf("at index %d, expected %s, Got %s.\n", i, test.out, value)
//         }
//     }
// }
