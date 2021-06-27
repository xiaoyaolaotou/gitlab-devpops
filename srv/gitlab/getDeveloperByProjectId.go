package gitlab

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"net/http"
)

// getDeveloperByProjectId 通过项目id获取开发者信息
func (g *GitLab) getDeveloperByProjectId (ctx context.Context, result Result)error {
	list,resp,err := g.git.ProjectMembers.ListAllProjectMembers(result.ProjectId,&gitlab.ListProjectMembersOptions{
		ListOptions:gitlab.ListOptions{
			Page: PageNum,
			PerPage: PageSize,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("获取失败")
	}
	for _,v := range list {
		result.UserName = v.Username
		result.AccessLevel = g.toString(v.AccessLevel)
		b,err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}
	return nil
}

// toString 将权限枚举类型转成string,方便阅读以及维护
func (g *GitLab) toString (AccessLevel gitlab.AccessLevelValue) string  {
	switch AccessLevel {
	case 0:
		return "NoPermissions"
	case 5:
		return "MinimalAccess"
	case 10:
		return "Guest"
	case 20:
		return "Reporter"
	case 30:
		return "Developer"
	case 40:
		return "Master"
	case 50:
		return "Owner"
	default:
		return "Unknown"
	}
}
