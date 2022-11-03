# quickmail
Temporary e-mail address generator library <br>
#### This project was made a while ago. I'm not sure if everything works
# Support
Currently only 1secmail is supported.

# Use example
```go
package main

import (
  "github.com/nextu1337/quickmail"
  "fmt"
  "time"
)

func main() {
  m := quickmail.New1SecMail();
  fmt.Println(m.GetMail());
  for ;; {
		msgs := m.GetMessages();
		for _,e := range msgs {
			fmt.Println(e.GetTitle());
			fmt.Println(e.GetBody());
		}
		time.Sleep(8*time.Second);
	}
}
```

# Objects
## Mail
- GetMail() string - returns e-mail address
- GetLogin() string - returns e-mail address' part before the @
- GetDomain() string - returns e-mail address' part after the @
- GetMessages() []*Message - returns mails sent to the address
## Message
- GetFrom() string - returns sender's e-mail address
- GetId() string - returns mail's id
- GetTitle() string - returns mail's title
- GetBody() string - returns mail's body
- GetDate() string - returns mail sent date (in string because i was too lazy to parse it to int)
