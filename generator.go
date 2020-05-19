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
func GenMultiLineMessage(width int, message string) (string, error) {
	messages, maxLength := getMessagesAndLength(message)
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
