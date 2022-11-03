package quickmail

type Message struct {
	id 		string
	from 	string
	title 	string
	body 	string
	date 	string
}

func (m *Message) GetFrom() string {
	return m.from
}

func (m *Message) GetId() string {
	return m.id
}

func (m *Message) GetTitle() string {
	return m.title
}

func (m *Message) GetBody() string {
	return m.body
}

func (m *Message) GetDate() string {
	return m.date
}





func NewMessage(id, from, title, body, date string) *Message {
	msg := new(Message)
	msg.id = id
	msg.from = from
	msg.title = title
	msg.body = body
	msg.date = date
	return msg
}