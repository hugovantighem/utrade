package app

type ApplicationError struct {
	Message string
	// TODO wrappable error
}

func (x ApplicationError) Error() string {
	return x.Message
}
