package todo

type Repository interface {
	Get(id string) (*Todo, error)
	Put(t *Todo) error
}
