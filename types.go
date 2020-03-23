package ted

import (
	"strings"
)

type Update struct {
	ID            int            `json:"update_id"`
	Message       *Message       `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
	InlineQuery   *InlineQuery   `json:"inline_query"`
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

type User struct {
	ID        int    `json:"ID"`
	FirstName string `json:"first_name"`
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
	ID string `json:"id"`

	// The user that chose the result
	From User `json:"from"`

	// Optional. Sender location, only for bots that require user location
	Location Location `json:"location"`

	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	InlineMessageID string `json:"inline_message_id"`

	// The query that was used to obtain the result
	Query string `json:"query"`
}
