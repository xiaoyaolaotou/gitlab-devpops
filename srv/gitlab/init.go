package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

// NewGitLabClient  实力化gitlab客户端
func NewGitLabClient () *gitlab.Client {
	git, err := gitlab.NewClient(Secret,gitlab.WithBaseURL(GitlabPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return git
}
