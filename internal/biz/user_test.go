package biz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("abc")
	fmt.Println(s)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$Jfe5dJqxf29LSH8M3.Nx0.iUzqaKzOYoya5pGn01Tu3mnaJOgwUEy", "abcl"))
}
