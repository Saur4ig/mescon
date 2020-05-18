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
		messageLength: utf8.RuneCountInString(message),
	}
	if width <= sm.messageLength+2 {
		return "", fmt.Errorf("message length more, than width %d > %d", sm.messageLength, sm.width)
	}

	if isMultiline(sm.message) {
		return "", fmt.Errorf("multiline message not allowed")
	}

	return sm.generateSingleLineMessage(), nil
}

// generates single line message
func (sm singleMessage) generateSingleLineMessage() string {
	var sb strings.Builder
	sb.WriteString(generateFullLine(sm.width))
	sb.WriteString(generateHollowLine(sm.width))
	sb.WriteString(wrapMessage(sm.width, sm.messageLength, sm.message))
	sb.WriteString(generateHollowLine(sm.width))
	sb.WriteString(generateFullLine(sm.width))
	return sb.String()
}

// wraps message with " *"
func wrapMessage(width, messageLen int, message string) string {
	sideAdder := (width - messageLen - 2) / 2
	for i := 0; i < sideAdder; i++ {
		message = " " + message + " "
	}
	return "*" + message + "*"
}

// returns string full of "*"
func generateFullLine(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteString("*")
	}
	return sb.String()
}

// returns string full of "*" with a new line in the end
func generateFullLineWithNL(length int) string {
	return generateFullLine(length) + "\n"
}

// returns string, that starts and ends with "*", and " " between them
func generateHollowLine(length int) string {
	var sb strings.Builder
	sb.WriteString("*")
	for i := 0; i < length-2; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("*")
	return sb.String()
}

// returns string, that starts and ends with "*", and " " between them with a new line in the end
func generateHollowLineWithNL(length int) string {
	return generateHollowLine(length) + "\n"
}

// is string contains new line symbol
func isMultiline(str string) bool {
	return strings.Contains(str, "\n")
}
