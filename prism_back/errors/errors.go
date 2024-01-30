package errors

import "fmt"

var (
	IsNotAdminUser error = fmt.Errorf("IsNotAdminUser")
)