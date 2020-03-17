package ted

import (
	"encoding/json"
)

type SendMessageRequest struct {
	ChatID      int         `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (r SendMessageRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("sendMessage", r)
}

type ReplyMarkup interface {
	replyMarkup()
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i InlineKeyboardMarkup) replyMarkup() {}

func (i InlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(data))
}

type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       string `json:"cache_time,omitempty"`
}

func (r AnswerCallbackQueryRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("answerCallbackQuery", r)
}
