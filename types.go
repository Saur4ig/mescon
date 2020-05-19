package mescon

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
