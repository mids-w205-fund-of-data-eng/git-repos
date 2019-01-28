// list command
package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
)

type Repo struct {
	Name    string
	OrgName string

	gh *GithubConnection
}

func NewRepo(orgName, repoName string) *Repo {

	gh, err := NewGithubConnection()
	if err != nil {
		log.Fatalf("GithubConnection: %s", err)
	}

	return &Repo{repoName, orgName, gh}
}

func (r *Repo) closedPRs() []*github.PullRequest {
	options := &github.PullRequestListOptions{
		State:       "closed",
		ListOptions: github.ListOptions{PerPage: 100},
	}
	pulls, _, err := r.gh.Client.PullRequests.List(r.gh.Context, r.OrgName, r.Name, options)
	if err != nil {
		log.Fatal(err)
	}
	return pulls
}

func (r *Repo) reviews(pr int) []*github.PullRequestReview {
	options := &github.ListOptions{PerPage: 100}
	reviews, _, err := r.gh.Client.PullRequests.ListReviews(r.gh.Context, r.OrgName, r.Name, pr, options)
	if err != nil {
		log.Fatal(err)
	}
	return reviews
}

func (r *Repo) trimSpaces(review string) string {
	return strings.Join(strings.Fields(review), " ")
}

func (r *Repo) ApprovedReviewBody() string {
	for _, pull := range r.closedPRs() {
		for _, review := range r.reviews(*pull.Number) {
			if review.User != nil && *review.User.Login == "mmm" && *review.State == "APPROVED" {
				return r.trimSpaces(*review.Body)
			}
		}
	}
	return ""
}

func (r *Repo) Grade() int {
	review := r.ApprovedReviewBody()
	justTheGrade := strings.Split(review, " ")[0]
	if justTheGrade != "" {
		grade, err := strconv.Atoi(justTheGrade[1:])
		if err != nil {
			log.Fatal(err)
		}
		return grade
	}
	return 0
}

func (r *Repo) NumClosedPRs() int {
	return len(r.closedPRs())
}
