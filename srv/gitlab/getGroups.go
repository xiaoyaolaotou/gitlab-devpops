package gitlab

import (
	"context"
	"errors"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"net/http"
)

// GetGroups 获取gitlab下面的所有组
func (g *GitLab) GetGroups (ctx context.Context) {
	var NextPageNum = PageNum
	var totalGroup = []*gitlab.Group{}
	list,err := g.groups(ctx,NextPageNum,PageSize)
	if err != nil {
		fmt.Println(err.Error())
	}
	totalGroup = append(totalGroup,list...)
	for { // 轮循到最后一页，为了解决gitlab接口支持全量获取数据的问题
		if len(list) == PageSize { // 获取的数据长度小于定义的页码。
			NextPageNum = NextPageNum + PageSize
			list,err = g.groups(ctx,NextPageNum,PageSize)
			if err != nil {
				fmt.Println(err.Error())
			}
			totalGroup = append(totalGroup,list...)
		}else { // 如果已经是最后一页直接跳出循环
			break
		}
	}
	// 最终取到的就是所有的组
	for _,v := range totalGroup {
		result := Result{
			GroupId: v.ID,
			GroupName: v.Name,
		}
		g.GetProjectByGroupId(ctx,result)
	}
	return
}

func (g *GitLab) groups (ctx context.Context,pageNum,pageSize int)([]*gitlab.Group,error) {
	list,resp,err := g.git.Groups.ListGroups(&gitlab.ListGroupsOptions{
		ListOptions:gitlab.ListOptions{
			Page: pageNum,
			PerPage: pageSize,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("获取当前页数据失败，httpCode = %d",resp.StatusCode))
	}
	return list,nil
}
