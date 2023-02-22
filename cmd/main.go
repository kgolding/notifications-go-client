package main

import (
	"flag"
	"fmt"
	"os"

	notificationclient "github.com/kgolding/notifications-go-client"
)

func main() {
	var APIKey, ServiceID string
	var Tel, TemplateID string

	flag.StringVar(&APIKey, "key", "", "API Key")
	flag.StringVar(&ServiceID, "id", "", "Service ID")
	flag.StringVar(&Tel, "tel", "", "Telephone number to send sms")
	flag.StringVar(&TemplateID, "tid", "", "Template ID")

	flag.Parse()

	if APIKey == "" {
		fmt.Println("missing key")
		os.Exit(1)
	}

	if ServiceID == "" {
		fmt.Println("missing id")
		os.Exit(1)
	}

	conf := notificationclient.Configuration{
		APIKey:    []byte(APIKey),
		ServiceID: ServiceID,
	}
	client, err := notificationclient.New(conf)
	if err != nil {
		fmt.Println("error creating new client:", err)
		os.Exit(1)
	}

	if Tel != "" {
		data := map[string]string{
			"message": "The message",
		}
		reference := ""
		r, err := client.SendSms(Tel, TemplateID, data, reference)
		if err != nil {
			fmt.Println("error sending sms:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("SMS Request success:\nContent: %s\nURI: %s\nID: %s\n",
			r.Content, r.URI, r.ID)
	}
}
