package errors

import "fmt"

var (
	IsNotAdminUser error = fmt.Errorf("IsNotAdminUser")
	NotSavedUser error = fmt.Errorf("NotSavedUser")
	EmptyFile error = fmt.Errorf("Empty File")
)