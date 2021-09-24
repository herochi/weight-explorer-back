package dto

type UserPatch struct {
	Username string `bson:"username" json:"username"`
	IsActive bool   `bson:"isActive" json:"isActive"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role" json:"role"`
}
