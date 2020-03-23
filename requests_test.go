package ted

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForceReply_MarshalJSON(t *testing.T) {
	var JSON []byte
	var err error
	JSON, err = json.Marshal(ForceReply{})
	assert.NoError(t, err)
	assert.Equal(t, `{"force_reply":true}`, string(JSON))
	JSON, err = json.Marshal(ForceReply{
		Selective: true,
	})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"force_reply":true,"selective":true}`, string(JSON))
}

func TestReplyKeyboardRemove_MarshalJSON(t *testing.T) {
	var JSON []byte
	var err error
	JSON, err = json.Marshal(ReplyKeyboardRemove{})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"remove_keyboard":true}`, string(JSON))
	JSON, err = json.Marshal(ReplyKeyboardRemove{
		Selective: true,
	})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"remove_keyboard":true,"selective":true}`, string(JSON))
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

func TestSendMessageRequest_MarshalJSON(t *testing.T) {
	req := SendMessageRequest{
		ChatID:    123,
		Text:      "Some text",
		ParseMode: "HTML",
		ReplyMarkup: InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{
						Text:         "Button",
						CallbackData: "Data",
					},
				},
			},
		},
	}
	expected := `{
  "chat_id": 123,
  "text": "Some text",
  "parse_mode": "HTML",
  "reply_markup": "{\"inline_keyboard\":[[{\"text\":\"Button\",\"callback_data\":\"Data\"}]]}"
}`
	actual, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func TestAnswerInlineQueryRequest_MarshalJSON(t *testing.T) {
	req := AnswerInlineQueryRequest{
		InlineQueryID: "123",
		Results: []InlineQueryResult{
			InlineQueryResultArticle{
				ID:    "456",
				Title: "Title",
				InputMessageContent: InputTextMessageContent{
					Text: "Text",
				},
				ReplyMarkup: &InlineKeyboardMarkup{
					InlineKeyboard: [][]InlineKeyboardButton{
						{
							{
								Text:         "Button",
								CallbackData: "Data",
							},
						},
					},
				},
				Description: "Description",
			},
		},
	}
	expected := `{
  "inline_query_id": "123",
  "results": "[{\"type\":\"article\",\"id\":\"456\",\"title\":\"Title\",\"input_message_content\":{\"message_text\":\"Text\"},\"reply_markup\":{\"inline_keyboard\":[[{\"text\":\"Button\",\"callback_data\":\"Data\"}]]},\"description\":\"Description\"}]"
}`
	actual, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func TestEditMessageReplyMarkupRequest_MarshalJSON(t *testing.T) {
	req := EditMessageReplyMarkupRequest{
		ChatID:          123,
		MessageID:       456,
		InlineMessageID: "abc",
		ReplyMarkup: &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{
						Text:         "Button",
						CallbackData: "Data",
					},
				},
			},
		},
	}
	expected := `{
  "chat_id": 123,
  "message_id": 456,
  "inline_message_id": "abc",
  "reply_markup": "{\"inline_keyboard\":[[{\"text\":\"Button\",\"callback_data\":\"Data\"}]]}"
}`
	actual, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func TestEditMessageTextRequest_MarshalJSON(t *testing.T) {
	req := EditMessageTextRequest{
		ChatID:          123,
		MessageID:       456,
		InlineMessageID: "abc",
		Text:            "New text",
		ReplyMarkup: &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{
						Text:         "Button",
						CallbackData: "Data",
					},
				},
			},
		},
	}
	expected := `{
  "chat_id": 123,
  "message_id": 456,
  "inline_message_id": "abc",
  "text": "New text",
  "reply_markup": "{\"inline_keyboard\":[[{\"text\":\"Button\",\"callback_data\":\"Data\"}]]}"
}`
	actual, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(actual))
}
