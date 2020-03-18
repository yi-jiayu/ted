package ted

import (
	"encoding/json"
)

type SendMessageRequest struct {
	// ChatID is a string when it refers to the username of a channel and an integer otherwise.
	// Required if InlineMessageID is not specified.
	ChatID interface{} `json:"chat_id"`

	// Text of the message to be sent. It should be limited to 1-4096 characters after entities parsing.
	Text string `json:"text"`

	// ParseMode can be specified to show bold, italic, fixed-width text or inline URLs in messages.
	// Possible values are "Markdown", "MarkdownV2" and "HTML". Refer to
	// https://core.telegram.org/bots/api#formatting-options for more information.
	ParseMode string `json:"parse_mode,omitempty"`

	// DisableWebPagePreview will disable link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// DisableNotification will send the message silently if set. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToMessageID is the ID of the message to reply to.
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// ReplyMarkup can be an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (r SendMessageRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("sendMessage", r)
}

type ReplyMarkup interface {
	replyMarkup()
}

// Upon receiving a message with this object, Telegram clients will display a
// reply interface to the user (act as if the user has selected the bot‘s
// message and tapped ’Reply'). This can be extremely useful if you want to
// create user-friendly step-by-step interfaces without having to sacrifice
// privacy mode.
type ForceReply struct {
	// Use this parameter if you want to force reply from specific users
	// only. Targets: 1) users that are @mentioned in the text of the
	// Message object; 2) if the bot's message is a reply (has
	// reply_to_message_id), sender of the original message.
	Selective bool
}

func (f ForceReply) replyMarkup() {}

func (f ForceReply) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(struct {
		ForceReply bool `json:"force_reply"`
		Selective  bool `json:"selective,omitempty"`
	}{
		ForceReply: true,
		Selective:  f.Selective,
	})
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(data))
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i InlineKeyboardMarkup) replyMarkup() {}

func (i InlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(struct {
		InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	}(i))
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

type EditMessageTextRequest struct {
	// ChatID is a string when it refers to the username of a channel and an integer otherwise.
	// Required if InlineMessageID is not specified.
	ChatID interface{} `json:"chat_id,omitempty"`

	// MessageID is required when InlineMessageID is not specified.
	MessageID int `json:"message_id,omitempty"`

	// InlineMessageID is required when ChatID and MessageID are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Text is the new text of the message. It should be limited to 1-4096 characters after entities parsing.
	Text string `json:"text"`

	// ParseMode can be specified to show bold, italic, fixed-width text or inline URLs in messages.
	// Possible values are "Markdown", "MarkdownV2" and "HTML". Refer to
	// https://core.telegram.org/bots/api#formatting-options for more information.
	ParseMode string `json:"parse_mode,omitempty"`

	// DisableWebPagePreview will disable link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// ReplyMarkup can be provided to display an inline keyboard with the updated message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageTextRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("editMessageText", e)
}
