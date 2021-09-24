package auth

type UserData struct {
	Username string `json:"username"`
	Role     string `bson:"role" json:"role"`
	User     string `bson:"user" json:"user"`
}
