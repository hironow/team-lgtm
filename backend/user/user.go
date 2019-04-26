package user

type User struct {
	ID string
}

func NewUser() *User {
	return &User{
		ID: "dummy user id",
	}
}
