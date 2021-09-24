package viewmodel

import "time"

type UserVM struct {
	ID        string     `bson:"_id,omitempty" json:"id"`
	IsActive  bool       `bson:"isActive" json:"isActive"`
	Username  string     `bson:"username" json:"username"`
	Password  string     `bson:"password" json:"password"`
	Role      string     `bson:"role" json:"role"`
	CreatedAt *time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt" json:"updatedAt"`
}
