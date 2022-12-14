package notification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnError_WhenToAdd_ThenShouldReturnSeveralErrorMessagesInAString(t *testing.T) {
	notification := NewSignUpNotification()
	notification.AddError("invalid phone number")
	notification.AddError("invalid email")
	notification.AddError("invalid password")
	assert.Equal(t, true, notification.HasErrors())
	assert.Equal(t, "signup: invalid phone number, invalid email, invalid password", notification.Messages())
}
