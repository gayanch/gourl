package analytics

import (
	"strconv"
)

type InMemoryAnalytics struct {
	data map[string]string
}

func (ima InMemoryAnalytics) Get(shortcode string) string {
	if val, ok := ima.data[shortcode]; ok {
		return val
	}
	return "Value not found"
}

func (ima InMemoryAnalytics) Update(shortcode string) {
	if val, ok := ima.data[shortcode]; ok {
		count, _ := strconv.Atoi(val);
		count += 1
		ima.data[shortcode] = string(count)
	} else {
		ima.data[shortcode] = "1"
	}
}