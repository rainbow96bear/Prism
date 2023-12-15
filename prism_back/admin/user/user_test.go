package User

import (
	"fmt"
	"testing"
)

func Test_HashPassword(t *testing.T) {
	password := "000000"

	result, err := HashPassword(password)
	if err != nil {
		t.Fatal("실패")
	}
	fmt.Println(result)
}

func Test_ComparePassword(t *testing.T) {
	password := "$2a$10$BB/MqSNa28rgfcon1./J.eh/KTgAlWKn6dO5D/GhFYfnb5Nb/ba0q"

	result := ComparePassword(password, "000000")
	fmt.Println(result)
}