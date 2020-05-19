package mescon

import (
	"strings"
)

// generates multi line message
func (mm multiLineMessage) generateMultiLineMessage() string {
	var sb strings.Builder
	sb.WriteString(generateFullLineWithNL(mm.width))
	sb.WriteString(generateHollowLineWithNL(mm.width))
	for _, val := range mm.messages {
		sb.WriteString(wrapMessage(mm.width, mm.maxMessageLength, val))
		sb.WriteString("\n")
	}
	sb.WriteString(generateHollowLineWithNL(mm.width))
	sb.WriteString(generateFullLineWithNL(mm.width))
	return sb.String()
}

// get messages as slice and max length inside
func getMessagesAndLength(message string) ([]string, int) {
	messages := getMessagesFromString(message)
	return messages, getMessageLengthInMessages(messages)
}

// get separate messages from one string
func getMessagesFromString(mes string) []string {
	return strings.Split(mes, "\n")
}

// get max length of message
func getMessageLengthInMessages(m []string) int {
	var max int
	for _, val := range m {
		if max < getMessageLength(val) {
			max = getMessageLength(val)
		}
	}
	return max
}
