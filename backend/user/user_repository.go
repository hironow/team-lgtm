package user

type Repository interface {
	Get(id string) (*User, error)
	Put(t *User) error
}
