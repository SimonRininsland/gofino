package model

// User
type User struct {
	Name string `json:"name" form:"name" query:"name"`
}

// AnomUser
type AnomUser []string
