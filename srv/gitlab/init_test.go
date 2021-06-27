package gitlab

import (
	"context"
	"testing"
)

func TestNewGitLabClient(t *testing.T) {
	client := GitLab{
		git: NewGitLabClient(),
	}
	ctx := context.Background()
	client.GetGroups(ctx)
}
