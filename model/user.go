package model

type (
	UserSignIn struct {
		Username	string
		Password	string
	}

	UserInfo struct {
		UserID		uint32	`json:"user_id"`
		Name		string	`json:"name"`
	}


)