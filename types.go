package mescon

// maximum message width
const MaxWidth = 125

// padding on the sides of the message
const SideSpaces = 3

// single line message
type singleMessage struct {
	width         int
	messageLength int
	message       string
}

type multiLineMessage struct {
	width            int
	maxMessageLength int
	messages         []string
}
