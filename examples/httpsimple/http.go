package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    []byte
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

func readRequest(reader *bufio.Reader) (*Request, error) {
	// read request line, e.g., "GET / HTTP/1.1"
	line, err := readLine(reader)
	if err != nil {
		return nil, fmt.Errorf("read status line: %w", err)
	}
	var method, path, httpVersion string
	_, err = fmt.Sscanf(line, "%s %s %s", &method, &path, &httpVersion)
	if err != nil {
		return nil, fmt.Errorf("parse request line: %w", err)
	}

	// read headers
	headers, err := readHeaders(reader)
	if err != nil {
		return nil, fmt.Errorf("read headers: %w", err)
	}

	return &Request{
		Method:  method,
		Path:    path,
		Headers: headers,
		Body:    []byte{},
	}, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("read line: %w", err)
	}

	// Note that TrimSpace removes '\r' as well.
	return strings.TrimSpace(line), nil
}

func readHeaders(reader *bufio.Reader) (map[string]string, error) {
	headers := make(map[string]string)
	for {
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		line = string(line)
		if line == "" {
			// end of header
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header format: %s", line)
		}
		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return headers, nil
}

func writeResponse(writer *bufio.Writer, res *Response) error {
	_, err := fmt.Fprintf(
		writer, "HTTP/1.1 %d %s\r\n",
		res.StatusCode,
		mapStatusCode(res.StatusCode),
	)
	if err != nil {
		return fmt.Errorf("write status line: %w", err)
	}

	for key, value := range res.Headers {
		_, err := fmt.Fprintf(writer, "%s: %s\r\n", key, value)
		if err != nil {
			return fmt.Errorf("write header: %w", err)
		}
	}

	_, err = writer.WriteString("\r\n")
	if err != nil {
		return fmt.Errorf("write header-body separator: %w", err)
	}

	_, err = writer.Write(res.Body)
	if err != nil {
		return fmt.Errorf("write body: %w", err)
	}

	return writer.Flush()
}

func mapStatusCode(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return ""
	}
}
