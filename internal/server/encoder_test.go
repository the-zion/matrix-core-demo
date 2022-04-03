package server

import (
	"cube-core/internal/errors"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTTPStruct(t *testing.T) {
	a := &errors.HTTPError{
		Errors: map[string][]string{},
	}
	a.Errors["body"] = []string{"can't be empty"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", b)
}
