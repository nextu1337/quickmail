package quickmail

type Provider int

const (
	ONESECMAIL Provider = iota
)

func GetBaseURL(provider Provider) string {
	return []string{"https://www.1secmail.com/api/v1/"}[provider]
}

type Mail struct {
	provider 	Provider
	mail 		string
	login 		string
	domain 		string
}

func (m *Mail) GetMail() string {
	return m.mail
}

func (m *Mail) GetLogin() string {
	return m.login
}

func (m *Mail) GetDomain() string {
	return m.domain
}

func (m *Mail) GetMessages() []*Message {
	var messages []*Message
	switch m.provider {
		case ONESECMAIL:
			messages = GetMessages1SecMail(m)
		break;
	}
	return messages
}