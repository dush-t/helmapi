package client

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// UpdateRepos is the equivalent of calling helm repo update
func UpdateRepos() error {
	app := "helm"
	cmd := exec.Command(app, "repo", "update")
	if err := execute(cmd); err != nil {
		return err
	}

	log.Println("Repos updated successfully")
	return nil
}

// RepoAddRequest represents a helm repo add command
type RepoAddRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Describe returns a string description (for printing) of
// the repo to be added
func (ra *RepoAddRequest) String() string {
	return fmt.Sprintf(`
--------------------
Name: %s
URL:  %s
--------------------
	`, ra.Name, ra.URL)
}

// Execute will add the repo as specified by RepoAddRequest
func (ra *RepoAddRequest) Execute() error {
	app := "helm"
	args := []string{"repo", "add", ra.Name, ra.URL}

	if len(ra.Name) == 0 || len(ra.URL) == 0 {
		return fmt.Errorf("URL or repo name cannot be empty")
	}

	log.Println("Adding repository:")
	log.Println(ra.String())

	cmd := exec.Command(app, args...)
	if err := execute(cmd); err != nil {
		return err
	}

	return nil
}

// RepoRemoveRequest represents a helm repo remove command
type RepoRemoveRequest struct {
	Repos []string `json:"repos"`
}

// Describe returns a string description (for printing) of
// the RepoRemoveRequest
func (rr *RepoRemoveRequest) String() string {
	return fmt.Sprintf(`
--------------------
Repos: %s
--------------------
	`, strings.Join(rr.Repos, ", "))
}

// Execute will remove the repos specified in RepoRemoveRequest
func (rr *RepoRemoveRequest) Execute() error {
	app := "helm"
	args := []string{"repo", "remove"}
	args = append(args, rr.Repos...)

	if len(rr.Repos) == 0 {
		return fmt.Errorf("You cannot provide empty repo list")
	}

	log.Println("Removing repos:")
	log.Println(rr.String())

	cmd := exec.Command(app, args...)
	if err := execute(cmd); err != nil {
		return err
	}

	return nil
}
