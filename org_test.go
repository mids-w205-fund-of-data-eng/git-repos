package main

import (
	"testing"
)

func TestOrg_List(t *testing.T) {
	o := NewOrg("mids-w205-fund-of-data-eng")

	matchingRepos := o.GetRepos("course-content")
	if len(matchingRepos) != 1 {
		t.Fatalf("bad: %d", len(matchingRepos))
	}
}

func TestOrg_DeleteRepoByFullName(t *testing.T) {
	o := NewOrg("mids-w205-martin-mims")

	testRepo, err := o.CreateRepo("test-repo-8456")
	if err != nil {
		t.Fatal(err)
	}

	err = o.DeleteRepoByName(*testRepo.Name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrg_DeleteRepoByPattern(t *testing.T) {
	o := NewOrg("mids-w205-martin-mims")

	_, err := o.CreateRepo("test-repo-8503")
	if err != nil {
		t.Fatal(err)
	}
	_, err = o.CreateRepo("test-repo-8504")
	if err != nil {
		t.Fatal(err)
	}

	err = o.DeleteReposByPattern("test-repo-85")
	if err != nil {
		t.Fatal(err)
	}
}
