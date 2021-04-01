package model

type UserSignIn struct {
	Username string
	Password string
}

type UserInfo struct {
	ID uint32
	Name string
}

type VotedWork struct {
	WorkID uint64	`json:"work_id"`
	ImageUrl string	`json:"image_url"`
}