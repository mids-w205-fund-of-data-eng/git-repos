package main

import (
	"testing"

	"github.com/google/uuid"
)

const testOrg = "mids-w205-martin-mims"

func generateRepoName() string {
	return "test-repo-" + uuid.New().String()[1:5]
}

func TestOrg_List(t *testing.T) {
	o := NewOrg(testOrg)

	matchingRepos := o.GetReposMatching("course-content")
	if len(matchingRepos) != 1 {
		t.Fatalf("failed to properly list repos in test org: %s", o.Name)
	}
}

func TestOrg_Create(t *testing.T) {
	o := NewOrg(testOrg)

	name := generateRepoName()
	_, err := o.CreateRepo(name)
	if err != nil {
		t.Fatal(err)
	}

	matchingRepos := o.GetReposMatching(name)
	if len(matchingRepos) != 1 {
		t.Fatalf("failed to create repo %s", name)
	}

	err = o.DeleteRepoByName(name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrg_DeleteRepoByName(t *testing.T) {
	o := NewOrg(testOrg)

	testRepo, err := o.CreateRepo(generateRepoName())
	if err != nil {
		t.Fatal(err)
	}

	err = o.DeleteRepoByName(*testRepo.Name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrg_DeleteRepoByPattern(t *testing.T) {
	o := NewOrg(testOrg)

	_, err := o.CreateRepo("somepattern-" + generateRepoName())
	if err != nil {
		t.Fatal(err)
	}
	_, err = o.CreateRepo("somepattern-" + generateRepoName())
	if err != nil {
		t.Fatal(err)
	}

	err = o.DeleteReposByPattern("somepattern-*")
	if err != nil {
		t.Fatal(err)
	}
}
