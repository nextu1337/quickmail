package quickmail

import (
    "encoding/json"
	"io/ioutil"
   	"log"
	"strings"
	"fmt"
   	"net/http"
)

type onesecmsg struct {
	Id 		int  `json:"id"`
	From 	string `json:"from"`
	Subject string `json:"subject"`
	Date 	string `json:"date"`
}

type onesecattachment struct {
	Filename string `json:"filename"`
	ContentType string `json:"contentType"`
	Size int `json:"size"`
}

type onesecread struct {
	Id int `json:"id"`
	From string `json:"from"`
	Subject string `json:"subject"`
	Date string `json:"date"`
	Attachments []onesecattachment `json:"attachments"`
	Body string `json:"body"`
	TextBody string `json:"textBody"`
	HtmlBody string `json:"htmlBody"`
}

func GetMessages1SecMail(m *Mail) []*Message {
	resp, err := http.Get(GetBaseURL(ONESECMAIL)+"?action=getMessages&login="+m.login+"&domain="+m.domain)
	if err != nil {
	   log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var msgs []*Message
	var msgs2 []onesecmsg
	err = json.Unmarshal(body,&msgs2)
	if err != nil {
		log.Fatalln(err)
	}
	for _,v := range msgs2 {
		fmt.Println(v);
		resp, err := http.Get(GetBaseURL(ONESECMAIL)+"?action=readMessage&login="+m.login+"&domain="+m.domain+"&id="+fmt.Sprintf("%d",v.Id))
		if err != nil {
		log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var msg onesecread
		err = json.Unmarshal(body,&msg)
		msgs = append(msgs, NewMessage(fmt.Sprintf("%d",v.Id),msg.From,msg.Subject,msg.TextBody,msg.Date))
	}
	return msgs
}

func New1SecMail() *Mail {
	var arr []string
	resp, err := http.Get(GetBaseURL(ONESECMAIL)+"?action=genRandomMailbox")
	if err != nil {
	   log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body,&arr)
	if err != nil {
		log.Fatalln(err)
	}
	mail := new(Mail)
	mail.provider = ONESECMAIL
	mail.mail = arr[0]
	mail.login = strings.Split(arr[0],"@")[0];
	mail.domain = strings.Split(arr[0],"@")[1];
	return mail
}