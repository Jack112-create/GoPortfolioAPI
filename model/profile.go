package model

type GithubProfile struct {
	Login        string `json:"login"`
	Avatar       string `json:"avatar_url"`
	Repositories string `json:"repos_url"`
	URL          string `json:"url"`
}
