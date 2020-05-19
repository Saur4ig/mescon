package mescon

import (
	"strings"
)

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
