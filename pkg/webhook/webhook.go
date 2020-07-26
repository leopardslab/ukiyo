package webhook

import docker "github.com/fsouza/go-dockerclient"

type Response struct {
	RepoName  string
	Namespace string
	Name      string
	Tag       string
	PushedAt  int
}

func DockerPullImage(str Response) docker.PullImageOptions {
	return docker.PullImageOptions{
		Repository:        str.RepoName,
		Tag:               str.Tag,
		Platform:          "",
		Registry:          "",
		OutputStream:      nil,
		RawJSONStream:     false,
		InactivityTimeout: 0,
		Context:           nil,
	}
}
