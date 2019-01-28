package main

import (
	"testing"
)

const testOrg = "mids-w205-martin-mims"

func TestNewRepo(t *testing.T) {
	t.Errorf("not implemented")
}

func TestRepo_List(t *testing.T) {
	r := NewOrg(testOrg)

	matchingRepos := o.GetReposMatching("course-content")
	if len(matchingRepos) != 1 {
		t.Fatalf("failed to properly list repos in test org: %s", o.Name)
	}
}
