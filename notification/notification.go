package notification

type Message struct {
	Body string
}

type Channel interface {
	Send(Message) error
	Name() string
}
