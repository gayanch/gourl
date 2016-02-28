package manager

import "testing"

func TestGenerateUrl(t *testing.T) {
    u := GenerateUrl()
    if len(u) != MAX_URL_LEN {
        t.Fatalf("Generate failed", len(u))
    }
}
