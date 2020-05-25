package mescon

import (
	"strings"
)

// generates single line message
func (sm singleMessage) generateSingleLineMessage() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(generateFullLine(sm.width))
	sb.WriteString(generateHollowLine(sm.width))
	sb.WriteString(wrapMessage(sm.width, sm.message))
	sb.WriteString(generateHollowLine(sm.width))
	sb.WriteString(generateFullLine(sm.width))
	return sb.String()
}
