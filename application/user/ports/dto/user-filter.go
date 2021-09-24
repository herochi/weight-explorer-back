package dto

type UserFilter struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
