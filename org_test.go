package main

import (
	"testing"
)

func TestOrg_List(t *testing.T) {
	o := NewOrg("mids-w205-fund-of-data-eng")
	matchingRepos := *o.GetRepos("course-content")
	if len(matchingRepos) != 1 {
		t.Fatalf("bad: %d", len(matchingRepos))
	}
}

func TestOrg_Flush(t *testing.T) {
}
