package ted

import (
	"encoding/json"
)

type SendMessageRequest struct {
	// ChatID is a string when it refers to the username of a channel and an integer otherwise.
	// Required if InlineMessageID is not specified.
	ChatID interface{}

	// Text of the message to be sent. It should be limited to 1-4096 characters after entities parsing.
	Text string

	// ParseMode can be specified to show bold, italic, fixed-width text or inline URLs in messages.
	// Possible values are "Markdown", "MarkdownV2" and "HTML". Refer to
	// https://core.telegram.org/bots/api#formatting-options for more information.
	ParseMode string

	// DisableWebPagePreview will disable link previews for links in this message.
	DisableWebPagePreview bool

	// DisableNotification will send the message silently if set. Users will receive a notification with no sound.
	DisableNotification bool

	// ReplyToMessageID is the ID of the message to reply to.
	ReplyToMessageID int

	// ReplyMarkup can be an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (r SendMessageRequest) MarshalJSON() ([]byte, error) {
	req := struct {
		ChatID                interface{} `json:"chat_id"`
		Text                  string      `json:"text"`
		ParseMode             string      `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
		DisableNotification   bool        `json:"disable_notification,omitempty"`
		ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           string      `json:"reply_markup,omitempty"`
	}{
		ChatID:                r.ChatID,
		Text:                  r.Text,
		ParseMode:             r.ParseMode,
		DisableWebPagePreview: r.DisableWebPagePreview,
		DisableNotification:   r.DisableNotification,
		ReplyToMessageID:      r.ReplyToMessageID,
	}
	if r.ReplyMarkup != nil {
		markup, err := json.Marshal(r.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		req.ReplyMarkup = string(markup)
	}
	return json.Marshal(req)
}

func (r SendMessageRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("sendMessage", r)
}

type ReplyMarkup interface {
	replyMarkup()
}

// This object represents one button of the reply keyboard. For simple text
// buttons String can be used instead of this object to specify text of the
// button. Optional fields request_contact, request_location, and request_poll
// are mutually exclusive.
type KeyboardButton struct {
	// 	Text of the button. If none of the optional fields are used, it
	// 	will be sent as a message when the button is pressed
	Text string `json:"text"`

	// If True, the user's phone number will be sent as a contact when the
	// button is pressed. Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`

	// If True, the user's current location will be sent when the button is
	// pressed. Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`
}

type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton or string.
	Keyboard [][]interface{}

	// Requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of
	// buttons). Defaults to false, in which case the custom keyboard is
	// always of the same height as the app's standard keyboard.
	ResizeKeyboard bool

	// Requests clients to hide the keyboard as soon as it's been used. The
	// keyboard will still be available, but clients will automatically
	// display the usual letter-keyboard in the chat – the user can press a
	// special button in the input field to see the custom keyboard again.
	// Defaults to false.
	OneTimeKeyboard bool

	// Use this parameter if you want to show the keyboard to specific
	// users only. Targets: 1) users that are @mentioned in the text of the
	// Message object; 2) if the bot's message is a reply (has
	// reply_to_message_id), sender of the original message.
	//
	// Example: A user requests to change the bot‘s language, bot replies
	// to the request with a keyboard to select the new language. Other
	// users in the group don’t see the keyboard.
	Selective bool
}

func (r ReplyKeyboardMarkup) replyMarkup() {}

// Upon receiving a message with this object, Telegram clients will remove the
// current custom keyboard and display the default letter-keyboard. By default,
// custom keyboards are displayed until a new keyboard is sent by a bot. An
// exception is made for one-time keyboards that are hidden immediately after
// the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	// Use this parameter if you want to remove the keyboard for specific
	// users only. Targets: 1) users that are @mentioned in the text of the
	// Message object; 2) if the bot's message is a reply (has
	// reply_to_message_id), sender of the original message.
	//
	// Example: A user votes in a poll, bot returns confirmation message in
	// reply to the vote and removes the keyboard for that user, while
	// still showing the keyboard with poll options to users who haven't
	// voted yet.
	Selective bool
}

func (r ReplyKeyboardRemove) replyMarkup() {}

func (r ReplyKeyboardRemove) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RemoveKeyboard bool `json:"remove_keyboard"`
		Selective      bool `json:"selective,omitempty"`
	}{
		RemoveKeyboard: true,
		Selective:      r.Selective,
	})
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
	return json.Marshal(struct {
		ForceReply bool `json:"force_reply"`
		Selective  bool `json:"selective,omitempty"`
	}{
		ForceReply: true,
		Selective:  f.Selective,
	})
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

// Use this method to send answers to callback queries sent from inline
// keyboards. The answer will be displayed to the user as a notification at the
// top of the chat screen or as an alert. On success, True is returned.
type AnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be shown to
	// the user, 0-200 characters
	Text string `json:"text,omitempty"`

	// If true, an alert will be shown by the client instead of a
	// notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// 	URL that will be opened by the user's client. If you have
	// 	created a Game and accepted the conditions via @Botfather,
	// 	specify the URL that opens your game – note that this will only
	// 	work if the query comes from a callback_game button.
	//
	// Otherwise, you may use links like t.me/your_bot?start=XXXX that open
	// your bot with a parameter.
	URL string `json:"url,omitempty"`

	// 	The maximum amount of time in seconds that the result of the
	// 	callback query may be cached client-side. Telegram apps will
	// 	support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
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

func (e EditMessageTextRequest) MarshalJSON() ([]byte, error) {
	req := struct {
		ChatID                interface{} `json:"chat_id,omitempty"`
		MessageID             int         `json:"message_id,omitempty"`
		InlineMessageID       string      `json:"inline_message_id,omitempty"`
		Text                  string      `json:"text"`
		ParseMode             string      `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
		ReplyMarkup           string      `json:"reply_markup,omitempty"`
	}{
		ChatID:                e.ChatID,
		MessageID:             e.MessageID,
		InlineMessageID:       e.InlineMessageID,
		Text:                  e.Text,
		ParseMode:             e.ParseMode,
		DisableWebPagePreview: e.DisableWebPagePreview,
	}
	if e.ReplyMarkup != nil {
		markup, err := json.Marshal(e.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		req.ReplyMarkup = string(markup)
	}
	return json.Marshal(req)
}

func (e EditMessageTextRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("editMessageText", e)
}

// This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 4 types:
//
//  InputTextMessageContent
//  InputLocationMessageContent
//  InputVenueMessageContent
//  InputContactMessageContent
type InputMessageContent interface {
	inputMessageContent()
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	Text string `json:"message_text"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

func (i InputTextMessageContent) inputMessageContent() {}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float32 `json:"latitude"`

	// Longitude of the location in degrees
	Longitude float32 `json:"longitude"`

	// Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
}

func (i InputLocationMessageContent) inputMessageContent() {}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	// Latitude of the location in degrees
	Latitude float32 `json:"latitude"`

	// Longitude of the location in degrees
	Longitude float32 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`
}

func (i InputVenueMessageContent) inputMessageContent() {}

// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
//
//  InlineQueryResultCachedAudio
//  InlineQueryResultCachedDocument
//  InlineQueryResultCachedGif
//  InlineQueryResultCachedMpeg4Gif
//  InlineQueryResultCachedPhoto
//  InlineQueryResultCachedSticker
//  InlineQueryResultCachedVideo
//  InlineQueryResultCachedVoice
//  InlineQueryResultArticle
//  InlineQueryResultAudio
//  InlineQueryResultContact
//  InlineQueryResultGame
//  InlineQueryResultDocument
//  InlineQueryResultGif
//  InlineQueryResultLocation
//  InlineQueryResultMpeg4Gif
//  InlineQueryResultPhoto
//  InlineQueryResultVenue
//  InlineQueryResultVideo
//  InlineQueryResultVoice
type InlineQueryResult interface {
	inlineQueryResult()
}

// InlineQueryResultArticle represents a link to an article or web page.
type InlineQueryResultArticle struct {
	// Unique identifier for this result, 1-64 Bytes
	ID string

	// Title of the result
	Title string

	// Content of the message to be sent
	InputMessageContent InputMessageContent

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup

	// Optional. URL of the result
	URL string

	// Optional. Pass True, if you don't want the URL to be shown in the message
	HideURL bool

	// Optional. Short description of the result
	Description string

	// Optional. Url of the thumbnail for the result
	ThumbURL string

	// Optional. Thumbnail width
	ThumbWidth int

	// Optional. Thumbnail height
	ThumbHeight int
}

func (i InlineQueryResultArticle) inlineQueryResult() {}

func (i InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		Title               string                `json:"title"`
		InputMessageContent InputMessageContent   `json:"input_message_content"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		URL                 string                `json:"url,omitempty"`
		HideURL             bool                  `json:"hide_url,omitempty"`
		Description         string                `json:"description,omitempty"`
		ThumbURL            string                `json:"thumb_url,omitempty"`
		ThumbWidth          int                   `json:"thumb_width,omitempty"`
		ThumbHeight         int                   `json:"thumb_height,omitempty"`
	}{
		Type:                "article",
		ID:                  i.ID,
		Title:               i.Title,
		InputMessageContent: i.InputMessageContent,
		ReplyMarkup:         i.ReplyMarkup,
		URL:                 i.URL,
		HideURL:             i.HideURL,
		Description:         i.Description,
		ThumbURL:            i.ThumbURL,
		ThumbWidth:          i.ThumbWidth,
		ThumbHeight:         i.ThumbHeight,
	})
}

// Use this method to send answers to an inline query. On success, True is
// returned. No more than 50 results per query are allowed.
type AnswerInlineQueryRequest struct {
	// Unique identifier for the answered query.
	InlineQueryID string

	// Results is an array of results for the inline query.
	Results []InlineQueryResult

	// The maximum amount of time in seconds that the result of the inline
	// query may be cached on the server. Defaults to 300.
	CacheTime int

	// Pass True, if results may be cached on the server side only for the
	// user that sent the query. By default, results may be returned to any
	// user who sends the same query
	IsPersonal bool

	// Pass the offset that a client should send in the next query with the
	// same text to receive more results. Pass an empty string if there are
	// no more results or if you don‘t support pagination. Offset length
	// can’t exceed 64 bytes.
	NextOffset string

	// If passed, clients will display a button with specified text that
	// switches the user to a private chat with the bot and sends the bot a
	// start message with the parameter switch_pm_parameter
	SwitchPMText string

	// Deep-linking parameter for the /start message sent to the bot when
	// user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9,
	// _ and - are allowed.
	//
	// Example: An inline bot that sends YouTube videos can ask the user to
	// connect the bot to their YouTube account to adapt search results
	// accordingly. To do this, it displays a ‘Connect your YouTube
	// account’ button above the results, or even before showing any. The
	// user presses the button, switches to a private chat with the bot
	// and, in doing so, passes a start parameter that instructs the bot to
	// return an oauth link. Once done, the bot can offer a switch_inline
	// button so that the user can easily return to the chat where they
	// wanted to use the bot's inline capabilities.
	SwitchPMParameter string
}

func (r AnswerInlineQueryRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("answerInlineQuery", r)
}

type InlineQueryResults []InlineQueryResult

func (r AnswerInlineQueryRequest) MarshalJSON() ([]byte, error) {
	req := struct {
		InlineQueryID     string `json:"inline_query_id"`
		Results           string `json:"results"`
		CacheTime         int    `json:"cache_time,omitempty"`
		IsPersonal        bool   `json:"is_personal,omitempty"`
		NextOffset        string `json:"next_offset,omitempty"`
		SwitchPMText      string `json:"switch_pm_text,omitempty"`
		SwitchPMParameter string `json:"switch_pm_parameter,omitempty"`
	}{
		InlineQueryID:     r.InlineQueryID,
		CacheTime:         r.CacheTime,
		IsPersonal:        r.IsPersonal,
		NextOffset:        r.NextOffset,
		SwitchPMText:      r.SwitchPMText,
		SwitchPMParameter: r.SwitchPMParameter,
	}
	data, err := json.Marshal(r.Results)
	if err != nil {
		return nil, err
	}
	req.Results = string(data)
	return json.Marshal(req)
}

// Use this method to edit only the reply markup of messages. On success, if
// edited message is sent by the bot, the edited Message is returned, otherwise
// True is returned.
type EditMessageReplyMarkupRequest struct {
	// Required if inline_message_id is not specified. Unique identifier
	// for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{}

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup
}

func (e EditMessageReplyMarkupRequest) MarshalJSON() ([]byte, error) {
	req := struct {
		ChatID          interface{} `json:"chat_id,omitempty"`
		MessageID       int         `json:"message_id,omitempty"`
		InlineMessageID string      `json:"inline_message_id,omitempty"`
		ReplyMarkup     string      `json:"reply_markup,omitempty"`
	}{
		ChatID:          e.ChatID,
		MessageID:       e.MessageID,
		InlineMessageID: e.InlineMessageID,
	}
	if e.ReplyMarkup != nil {
		markup, err := json.Marshal(e.ReplyMarkup)
		if err != nil {
			return nil, err
		}
		req.ReplyMarkup = string(markup)
	}
	return json.Marshal(req)
}

func (e EditMessageReplyMarkupRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("editMessageReplyMarkup", e)
}
