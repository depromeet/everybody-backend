package noti

type NotifierPort interface {
	Send(sender string, body *Message) error
}

type Message struct {
	Sender  string
	Title   string
	Content string
}
