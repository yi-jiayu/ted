package ted

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyKeyboardMarkup_MarshalJSON(t *testing.T) {
	replyKeyboard := ReplyKeyboardMarkup{
		Keyboard: [][]interface{}{
			{
				KeyboardButton{
					Text:           "Text",
					RequestContact: true,
				},
				"String",
			},
			{
				"Another string",
			},
		},
		ResizeKeyboard: true,
	}
	actual, err := json.Marshal(replyKeyboard)
	assert.NoError(t, err)
	assert.Equal(t, `"{\"keyboard\":[[{\"text\":\"Text\",\"request_contact\":true},\"String\"],[\"Another string\"]],\"resize_keyboard\":true}"`, string(actual))
}

func TestInlineKeyboardMarkup_MarshalJSON(t *testing.T) {
	markup := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{
					Text:         "Text",
					CallbackData: "Data",
				},
			},
		},
	}
	actual, err := json.Marshal(markup)
	assert.NoError(t, err)
	assert.Equal(t, `"{\"inline_keyboard\":[[{\"text\":\"Text\",\"callback_data\":\"Data\"}]]}"`, string(actual))
}

func TestForceReply_MarshalJSON(t *testing.T) {
	var JSON []byte
	var err error
	JSON, err = json.Marshal(ForceReply{})
	assert.NoError(t, err)
	assert.Equal(t, `"{\"force_reply\":true}"`, string(JSON))
	JSON, err = json.Marshal(ForceReply{
		Selective: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, `"{\"force_reply\":true,\"selective\":true}"`, string(JSON))
}

func TestReplyKeyboardRemove_MarshalJSON(t *testing.T) {
	var JSON []byte
	var err error
	JSON, err = json.Marshal(ReplyKeyboardRemove{})
	assert.NoError(t, err)
	assert.Equal(t, `"{\"remove_keyboard\":true}"`, string(JSON))
	JSON, err = json.Marshal(ReplyKeyboardRemove{
		Selective: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, `"{\"remove_keyboard\":true,\"selective\":true}"`, string(JSON))
}
