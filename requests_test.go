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

func TestInlineQueryResultArticle_MarshalJSON(t *testing.T) {
	result := InlineQueryResultArticle{
		ID:    "123",
		Title: "Title",
		InputMessageContent: InputTextMessageContent{
			Text: "Message text",
		},
		Description: "Description",
		ThumbURL:    "Thumbnail URL",
		ThumbWidth:  200,
		ThumbHeight: 200,
	}
	JSON, err := json.Marshal(result)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
  "type": "article",
  "id": "123",
  "title": "Title",
  "input_message_content": {
    "message_text": "Message text"
  },
  "description": "Description",
  "thumb_url": "Thumbnail URL",
  "thumb_width": 200,
  "thumb_height": 200
}`, string(JSON))
}
