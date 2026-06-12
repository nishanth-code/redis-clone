package server

import "strings"

type Command struct {
	Action string
	Key    string
	Value  string
}

func ParseCommand(input string) Command {
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)

	cmd := Command{}

	if len(parts) > 0 {
		cmd.Action = strings.ToUpper(parts[0])
	}

	if len(parts) > 1 {
		cmd.Key = parts[1]
	}

	if len(parts) > 2 {
		cmd.Value = parts[2]
	}

	return cmd
}