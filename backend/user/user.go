package user

type User struct {
	ID string
}

func NewUser() *User {
	return &User{
		ID: generateID(),
	}
}

func generateID() string {
	return "dummy user id"
}