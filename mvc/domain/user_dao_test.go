package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.ErrorCode, "They should be equal")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.NotNil(t, user, "User should not be nil")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Rishab", user.FirstName, "First Name of user should be Rishab")
}
