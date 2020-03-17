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
	assert.Equal(t, `"{\"inline_keyboard\":[[{\"text\":\"Text\",\"url\":\"\",\"callback_data\":\"Data\"}]]}"`, string(actual))
}
