package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		inNames, wantNames []string
	}{
		{[]string{"Xiong, Fong", "Ling, Mai"}, []string{"Ling, Mai", "Xiong, Fong"}},
	}

	for _, c := range cases {
		gotNames := Sort(c.inNames)
		for idx, name := range c.wantNames {
			if name != gotNames[idx] {
				t.Errorf("sortResult[%d] == %q, want %q", idx, gotNames[idx], name)
			}
		}
	}
}
