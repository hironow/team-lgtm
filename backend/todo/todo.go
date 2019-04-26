package todo

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

func NewTodo() *Todo {
	return &Todo{
		ID: "dummy todo id",
	}
}
