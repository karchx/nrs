package platform

import (
	log "github.com/gothew/l-og"
	"github.com/karchx/nrs/pkg/issue"
	"github.com/karchx/nrs/pkg/platform/github"
) 

func GetCredentials() []issue.IssuesAPI {
  creds := []issue.IssuesAPI{}

  if github, err := github.GetGithubCredentials(); err == nil {
    log.Info(github)
  }
  
  return creds
}
