package project

import "regexp"

// Project contains details of git repository
type Project struct {
	Protocol string
	Provider string
	Name     string
}

// GetProjectFromRemoteURL parses git repo remote URL into project.Project.
// We only accepts 2 formats for now:
// git@github.com:user/git-repo-name.git
// https://github.com/user/git-repo-name.git
func GetProjectFromRemoteURL(remoteURL string) *Project {
	re := regexp.MustCompile(`(?P<Protocol>git@|http(s)?:\/\/)(.+@)*(?P<Provider>[\w\d\.]+)(:[\d]+){0,1}\/*(?P<ProjectName>.*)`)
	matches := re.FindStringSubmatch(remoteURL)

	protocolRe := regexp.MustCompile(`[^\w]`)
	protocol := protocolRe.ReplaceAllString(matches[re.SubexpIndex("Protocol")], "")

	providerRe := regexp.MustCompile(`^(.*?)\.`)
	provider := providerRe.FindStringSubmatch(matches[re.SubexpIndex("Provider")])

	projectNameRe := regexp.MustCompile(`\/(.*)\.git$`)
	projectName := projectNameRe.FindStringSubmatch(matches[re.SubexpIndex("ProjectName")])

	return &Project{
		Protocol: protocol,
		Provider: provider[1],
		Name:     projectName[1],
	}
}
