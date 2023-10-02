package domain

type TodoStatus string

const (
	Created TodoStatus = "Created"
	Done    TodoStatus = "Done"
)

func (x TodoStatus) IsValid() bool {
	switch x {
	case Created, Done:
		return true
	default:
		return false
	}
}
