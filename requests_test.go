package ted

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
