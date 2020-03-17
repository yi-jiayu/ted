package ted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMessageNotModified(t *testing.T) {
	var err error = Response{
		OK:          false,
		ErrorCode:   400,
		Description: "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message",
	}
	assert.True(t, IsMessageNotModified(err))
}
