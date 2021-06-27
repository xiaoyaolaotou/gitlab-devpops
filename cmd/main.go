package main

import (
	"gitlab-devops/srv/gitlab"
	"net/http"
)

func main () {
	gitlab.Cron()
	http.ListenAndServe(":9090", nil)
}


