package main

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/PhilippReinke/tcp-to-http/pkg/protocol"
)

type Chat struct {
	users map[username]userInfo
}

func NewChat() *Chat {
	return &Chat{
		users: make(map[username]userInfo),
	}
}

type username = string

type userInfo struct {
	passwordHash string
}

func (c *Chat) RegisterUser(name, password string) error {
	nameLower := strings.ToLower(name)
	if nameLower == "system" {
		return errors.New("'system' user is reserved")
	}
	_, ok := c.users[nameLower]
	if ok {
		return fmt.Errorf("username already taken")
	}

	c.users[nameLower] = userInfo{
		passwordHash: string(md5.New().Sum([]byte(password))),
	}

	return nil
}

var _ protocol.Protocol = (*Chat)(nil)

func (c *Chat) HandleConnection(
	conn net.Conn,
	broadcaster protocol.Broadcaster,
) error {
	var username string
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// sync writes via channel
	writeChan := make(chan string)
	var writeErr error
	defer close(writeChan)
	go func() {
		for {
			var msgOut string
			select {
			case msg, ok := <-writeChan:
				if !ok {
					// channel has been closed
					return
				}
				msgOut = msg
			case msgByte := <-broadcaster.Send:
				if username == "" {
					// not logged in
					continue
				}
				msgOut = string(msgByte)
			}

			_, err := writer.WriteString(msgOut + "\r\n")
			if err != nil {
				writeErr = fmt.Errorf("writing string: %w", err)
				return
			}

			if err = writer.Flush(); err != nil {
				writeErr = fmt.Errorf("flushing: %w", err)
				return
			}
		}
	}()

	defer func() {
		if username != "" {
			sendBroadcaster(
				broadcaster.Receive,
				fmt.Sprintf("system %s has left", username),
			)
		}
	}()

	// read loop
	for {
		line, err := readLine(reader)
		if err != nil {
			sendSystemMessage(writeChan, err.Error())
			continue
		}

		command, err := parseCommand(line)
		if err != nil {
			sendSystemMessage(writeChan, err.Error())
			continue
		}

		switch command.Keyword {
		case REGISTER:
			if len(command.Args) != 2 {
				sendSystemMessage(
					writeChan,
					"usage: REGISTER <name> <password>",
				)
				continue
			}
			name := command.Args[0]
			password := command.Args[1]
			if err := c.RegisterUser(name, password); err != nil {
				sendSystemMessage(writeChan, err.Error())
				continue
			}
			sendSystemMessage(writeChan, fmt.Sprintf("registered %s", name))
		case LOGIN:
			if len(command.Args) != 2 {
				sendSystemMessage(writeChan, "usage: LOGIN <name> <password>")
				continue
			}
			name := command.Args[0]
			password := command.Args[1]

			user, ok := c.users[strings.ToLower(name)]
			if !ok ||
				user.passwordHash != string(md5.New().Sum([]byte(password))) {
				sendSystemMessage(writeChan, "invalid username or password")
				continue
			}

			username = name // set the username for the current connection
			sendSystemMessage(writeChan, fmt.Sprintf("logged in as %s", name))
		case LIST:
			var registeredUsers []string
			for name := range c.users {
				registeredUsers = append(registeredUsers, name)
			}
			sendSystemMessage(
				writeChan,
				fmt.Sprintf(
					"registered users: %s",
					strings.Join(registeredUsers, ", "),
				),
			)
		case MSG:
			if username == "" {
				sendSystemMessage(writeChan, "not logged in")
				continue
			}

			if len(command.Args) == 0 {
				sendSystemMessage(writeChan, "usage: MSG <message>")
				continue
			}

			message := strings.Join(command.Args, " ")
			sendBroadcaster(
				broadcaster.Receive,
				fmt.Sprintf("%s %s", username, message),
			)
		case PRIVMSG:
			// needs to be implemented
		case BYE:
			sendSystemMessage(writeChan, "goodbye")
			return writeErr
		default:
			sendSystemMessage(writeChan, "unknown command")
		}
	}
}

func sendSystemMessage(writeChan chan string, msg string) {
	writeChan <- fmt.Sprintf("system %v\r\n", msg)
}

func sendBroadcaster(writeChan chan<- []byte, msg string) {
	writeChan <- []byte(fmt.Sprintf("%v\r\n", msg))
}
