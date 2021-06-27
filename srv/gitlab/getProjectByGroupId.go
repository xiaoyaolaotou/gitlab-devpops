package gitlab

import (
	"context"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"net/http"
)

// GetProjectByGroupId 通过组id 获取项目列表
func (g *GitLab)GetProjectByGroupId (ctx context.Context,result Result) error {
	projects,resp,err := g.git.Groups.ListGroupProjects(result.GroupId,&gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page: PageNum,
			PerPage: PageSize,
		},
	})
	if err !=  nil {
		fmt.Println(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("请求失败")
	}
	for _,v:= range projects {
		result.ProjectId = v.ID
		result.ProjectName = v.Name
		g.getDeveloperByProjectId(ctx,result)
	}
	return nil
}
