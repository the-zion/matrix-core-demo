package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("hello", "eric")
	spew.Dump(tk)
}
