package gitlab

import (
	"context"
	"fmt"
	"time"
)


// Cron 定时任务
func Cron () {
	t := time.Tick(time.Second*10) // 十秒执行一次，按照需求改成自己想要的是假节点即可
	for {
		select {
		case <- t:
			// TODO 调用api更新gitlab仓库信息
			fmt.Println("tick------触发定时任务------")
			g := GitLab{
				git: NewGitLabClient(),
			}
			g.GetGroups(context.TODO())
		}
	}
}
