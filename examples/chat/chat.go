package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type CommandType int

const (
	REGISTER CommandType = iota
	LOGIN
	LIST
	MSG
	PRIVMSG
	BYE
)

type Command struct {
	Keyword CommandType
	Args    []string
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("read line: %w", err)
	}

	// Note that TrimSpace removes '\r' as well.
	return strings.TrimSpace(line), nil
}

func parseCommand(line string) (Command, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return Command{}, errors.New("empty command")
	}

	keywordStr := strings.ToUpper(parts[0])
	var commandType CommandType
	switch keywordStr {
	case "REGISTER":
		commandType = REGISTER
	case "LOGIN":
		commandType = LOGIN
	case "LIST":
		commandType = LIST
	case "MSG":
		commandType = MSG
	case "PRIVMSG":
		commandType = PRIVMSG
	case "BYE":
		commandType = BYE
	default:
		return Command{}, fmt.Errorf("unknown command: %s", keywordStr)
	}

	return Command{
		Keyword: commandType,
		Args:    parts[1:],
	}, nil
}
