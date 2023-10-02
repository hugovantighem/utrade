package domain

import "fmt"

type InvalidArgumentError struct {
	ArgName string
	Reason  string
}

func (x InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid %s: %s", x.ArgName, x.Reason)
}

type NotFoundError struct {
	Target      string
	LookupKey   string
	LookupValue string
}

func (x NotFoundError) Error() string {
	return fmt.Sprintf("%s not found (%s=%s)", x.Target, x.LookupKey, x.LookupValue)
}
