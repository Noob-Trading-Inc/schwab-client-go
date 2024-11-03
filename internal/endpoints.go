package internal

type endpoints struct {
	Auth  string
	Token string
}

var Endpoints = &endpoints{
	Auth:  "https://api.schwabapi.com/v1/oauth/authorize",
	Token: "https://api.schwabapi.com/v1/oauth/token",
}
