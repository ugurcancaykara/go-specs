package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	// Test verisi
	pages := map[string]int{
		"https://www.wagslane.dev":                               6,
		"https://www.wagslane.dev/about":                         5,
		"https://www.wagslane.dev/index.xml":                     5,
		"https://www.wagslane.dev/tags":                          5,
		"https://www.wagslane.dev/posts/go-struct-ordering":      1,
		"https://www.wagslane.dev/posts/seo-is-a-scam-job":       1,
		"https://www.wagslane.dev/posts/managers-that-cant-code": 1,
	}

	// Beklenen sıralı sonuç
	expected := []Page{
		{URL: "https://www.wagslane.dev", Count: 6},
		{URL: "https://www.wagslane.dev/about", Count: 5},
		{URL: "https://www.wagslane.dev/index.xml", Count: 5},
		{URL: "https://www.wagslane.dev/tags", Count: 5},
		{URL: "https://www.wagslane.dev/posts/go-struct-ordering", Count: 1},
		{URL: "https://www.wagslane.dev/posts/managers-that-cant-code", Count: 1},
		{URL: "https://www.wagslane.dev/posts/seo-is-a-scam-job", Count: 1},
	}

	// Fonksiyonu çalıştırma
	sortedPages := sortPages(pages)

	// Beklenenle karşılaştırma
	if !reflect.DeepEqual(sortedPages, expected) {
		t.Errorf("Got %v, expected %v", sortedPages, expected)
	}
}
