package domain

type Todo struct {
	msg    string
	userID string
	status TodoStatus
}

func NewTodo(msg string, userID string) (Todo, error) {
	if len(msg) == 0 {
		return Todo{}, InvalidArgumentError{
			ArgName: "msg",
			Reason:  "should not be empty",
		}
	}

	if len(userID) == 0 {
		return Todo{}, InvalidArgumentError{
			ArgName: "userID",
			Reason:  "should not be empty",
		}
	}

	return Todo{
		msg:    msg,
		userID: userID,
		status: Created,
	}, nil
}

// getters
func (x Todo) Message() string {
	return x.msg
}

func (x Todo) UserID() string {
	return x.userID
}

func (x Todo) Status() TodoStatus {
	return x.status
}
