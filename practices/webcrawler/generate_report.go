package main

import (
	"fmt"
	"sort"
)

// Sorting logic of pages
func sortPages(pages map[string]int) []Page {
	var pageSlice []Page
	for url, count := range pages {
		pageSlice = append(pageSlice, Page{URL: url, Count: count})
	}

	// sort first depending on the count then depending on the url
	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].Count == pageSlice[j].Count {
			return pageSlice[i].URL < pageSlice[j].URL
		}
		return pageSlice[i].Count > pageSlice[j].Count
	})
	return pageSlice
}

// purpose of this func is to sort the pages map and make it look good when it's printed
func generateReport(baseURL string, pages map[string]int) string {

	sortedPages := sortPages(pages)

	var report string
	report += fmt.Sprintf("=============================\n  REPORT for %s\n=============================\n", baseURL)
	for _, page := range sortedPages {
		report += fmt.Sprintf("Found %d internal links to %s\n", page.Count, page.URL)
	}
	return report
}
