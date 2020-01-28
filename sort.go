package main

import (
	"sort"
)

func Sort(names []string) []string {
	sort.Strings(names)
	return names
}
