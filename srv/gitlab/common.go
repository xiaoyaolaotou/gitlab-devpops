package gitlab

import (
	"context"

	"github.com/xanzy/go-gitlab"
)

const (
	PageNum    = 1
	PageSize   = 10
	GitlabPath = "http://47.97.223.93:81/api/v4"
	Secret     = "oFmo8xBTeKHLrWScrEhv"
)

type GitLab struct {
	git *gitlab.Client
}

type Interface interface {
	// 结口实现
	GetGroups(ctx context.Context)
	GetProjectByGroupId(ctx context.Context, result Result) error
	getDeveloperByProjectId(ctx context.Context, result Result) error
	groups(ctx context.Context, pageNum, pageSize int) ([]*gitlab.Group, error)
	toString(AccessLevel gitlab.AccessLevelValue) string
}

type Result struct {
	GroupId     int    `json:"groupId"`     // 组id
	GroupName   string `json:"groupName"`   // 组名称
	ProjectId   int    `json:"projectId"`   // 项目id
	ProjectName string `json:"projectName"` // 项目名称
	UserName    string `json:"userName"`    // 用户名称
	AccessLevel string `json:"accessLevel"` // 权限等级
}
