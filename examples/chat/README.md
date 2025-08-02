# Chat

**Implementation not finished**

The chat protocol is a line-based protocol with messages of the form:

```
COMMAND arg1 arg2\r\n
```

The following commands are supported:

|          command           |           description            |
| -------------------------- | -------------------------------- |
| REGISTER <name> <password> | register a new user              |
| LOGIN <name> <password>    | login as user                    |
| LIST                       | list registered users            |
| MSG <message>              | send messages to chat            |
| PRIVMSG <name> <msg>       | send a private message to a user |
| BYE                        | leave chat                       |

Note the following restrictions:

- usernammes must be lower-case
- username and password must not contain whitespaces
- `system` user is reserved for error messages

Responses are of the form:

```
<name> <message>\r\n
```

## Testing it

Note that there is a `client.html` that can be used to interact with the chat.
