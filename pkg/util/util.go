package util

type PullObj struct {
	RepoName   string
	Namespace  string
	Tag        string
	PushedDate int
}

type ResponseObj struct {
	ResponseCode int
	ResponseDesc string
}
