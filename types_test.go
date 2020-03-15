package ted

import (
	"testing"
)

func TestMessage_CommandAndArgs(t *testing.T) {
	tests := []struct {
		name        string
		message     Message
		wantCommand string
		wantArgs    string
	}{
		{
			name: "removes leading slash from command",
			message: Message{
				Text: "/eta",
				Entities: []MessageEntity{
					{
						Type:   "bot_command",
						Offset: 0,
						Length: 4,
					},
				},
			},
			wantCommand: "eta",
			wantArgs:    "",
		},
		{
			name: "trims whitespace from args",
			message: Message{
				Text: "/cmd hello",
				Entities: []MessageEntity{
					{
						Type:   "bot_command",
						Offset: 0,
						Length: 4,
					},
				},
			},
			wantCommand: "cmd",
			wantArgs:    "hello",
		},
		{
			name: "removes bot mention from command",
			message: Message{
				Text: "/cmd@username args",
				Entities: []MessageEntity{
					{
						Type:   "bot_command",
						Offset: 0,
						Length: 13,
					},
				},
			},
			wantCommand: "cmd",
			wantArgs:    "args",
		},
		{
			name: "returns message text as args when command not present",
			message: Message{
				Text: "args",
			},
			wantCommand: "",
			wantArgs:    "args",
		},
		{
			name: "ignores command not at beginning of message",
			message: Message{
				Text: "Try this command: /cmd",
				Entities: []MessageEntity{
					{
						Type:   "bot_command",
						Offset: 18,
						Length: 4,
					},
				},
			},
			wantCommand: "",
			wantArgs:    "Try this command: /cmd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command, args := tt.message.CommandAndArgs()
			if command != tt.wantCommand {
				t.Fatalf("command: want %q, got %q", tt.wantCommand, command)
			}
			if args != tt.wantArgs {
				t.Fatalf("args: want %q, got %q", tt.wantCommand, args)
			}
		})
	}
}
