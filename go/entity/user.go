package entity

// User is user
type User struct {
	Id   int
	Name string
}

func (User) TableName() string {
	return "users"
}
