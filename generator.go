package mescon

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// generates cosy message only for one line message
func GenSingleLineMessage(width int, message string) (string, error) {
	sm := singleMessage{
		width:         width,
		message:       message,
		messageLength: getMessageLength(message),
	}
	if width <= sm.messageLength+2 {
		return "", fmt.Errorf("message length more, than width %d > %d", sm.messageLength, sm.width)
	}

	if isMultiline(sm.message) {
		return "", fmt.Errorf("multiline message not allowed")
	}

	return sm.generateSingleLineMessage(), nil
}

// generates cosy message for multiple lines messages
func GenMultiLineMessage(width int, message, customSeparator string) (string, error) {
	separator := "\n"
	if customSeparator != "" {
		separator = customSeparator
	}
	messages, maxLength := getMessagesAndLength(message, separator)
	if maxLength+2 >= width {
		return "", fmt.Errorf("message length more, than width %d > %d", maxLength, width)
	}
	multi := multiLineMessage{
		width:            width,
		messages:         messages,
		maxMessageLength: maxLength,
	}
	return multi.generateMultiLineMessage(), nil
}

// generates lovely message with any amount of lines, and without fixed width
func GenAny(message string) (string, error) {
	messages, maxLength := getMessagesAndLength(message, "\n")
	width := maxLength + SideSpaces + 2
	if width > MaxWidth {
		return "", fmt.Errorf("message width more, than max width %d > %d", width, MaxWidth)
	}
	if len(messages) > 1 {
		mu := multiLineMessage{
			width:            maxLength + SideSpaces + 2,
			maxMessageLength: maxLength,
			messages:         messages,
		}
		return mu.generateMultiLineMessage(), nil
	}
	sm := singleMessage{
		width:         width,
		message:       messages[0],
		messageLength: maxLength,
	}
	return sm.generateSingleLineMessage(), nil
}

// wraps message with " *"
func wrapMessage(width, messageLen int, message string) string {
	sideAdder := (width - messageLen - 2) / 2
	for i := 0; i < sideAdder; i++ {
		message = " " + message + " "
	}
	if (sideAdder*2+messageLen)+2 != width {
		return "*" + message + " *"
	}
	return "*" + message + "*"
}

// is string contains new line symbol
func isMultiline(str string) bool {
	return strings.Contains(str, "\n")
}

// get amount of chars in messages
func getMessageLength(mes string) int {
	return utf8.RuneCountInString(mes)
}
