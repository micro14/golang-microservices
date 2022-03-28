package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user)
	assert.Equal(t, err.ErrorCode, http.StatusNotFound, "They should be equal")
}
