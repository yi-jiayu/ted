package ted

import (
	"strings"
)

type Update struct {
	ID                 int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
}

type Message struct {
	// Unique message identifier inside this chat
	ID int `json:"message_id"`

	// Sender, empty for messages sent to channels
	From *User `json:"from"`

	// Conversation the message belongs to
	Chat Chat `json:"chat"`

	// For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message"`

	// For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Text string `json:"text"`

	// For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location"`
}

// CommandAndArgs extracts and returns a Telegram bot command and the rest of
// the message excluding the command. The command will have its leading slash
// and possible bot mention removed. If a command was not present in the
// message or the message did not begin with a command, command will be an
// empty string and args will contain the entire message text.
func (m Message) CommandAndArgs() (string, string) {
	for _, e := range m.Entities {
		if e.Type == "bot_command" && e.Offset == 0 {
			command := strings.TrimPrefix(m.Text[:e.Length], "/")
			args := strings.TrimSpace(m.Text[e.Length:])
			mention := strings.Index(command, "@")
			if mention != -1 {
				return command[:mention], args
			}
			return command, args
		}
	}
	return "", m.Text
}

// User represents a Telegram user or bot.
type User struct {
	// Unique identifier for this user or bot
	ID int `json:"ID"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// User's or bot's first name
	FirstName string `json:"first_name"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name"`

	// Optional. User's or bot's username
	Username string `json:"username"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}

type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	Data            string   `json:"data"`
}

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
//
// Note: It is necessary to enable inline feedback via @Botfather in order to receive these objects in updates.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ID string `json:"result_id"`

	// The user that chose the result
	From User `json:"from"`

	// Optional. Sender location, only for bots that require user location
	Location Location `json:"location"`

	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	InlineMessageID string `json:"inline_message_id"`

	// The query that was used to obtain the result
	Query string `json:"query"`
}

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`

	// HasCustomCertificate will be true if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// PendingUpdateCount is the number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`

	// Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate int `json:"last_error_date"`

	// Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message"`

	// Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections int `json:"max_connections"`

	// Optional. A list of update types the bot is subscribed to. Defaults to all update types
	AllowedUpdates []string `json:"allowed_updates"`
}
