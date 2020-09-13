package models

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistration struct {
	UserId   string `json:"userId,omitempty" bson:"userId"`
	Username string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
}
