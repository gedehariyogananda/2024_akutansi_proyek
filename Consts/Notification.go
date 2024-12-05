package Consts

import "firebase.google.com/go/messaging"

// Status
const (
	Pending = "pending"
	Skipped = "skipped"
	Sent    = "sent"
)

type PushNotificationScheme string

// Scheme
const ()

func (i PushNotificationScheme) DefinePushNotificationMessages() *messaging.Notification {
	switch i {
	// todo :: add more scheme
	case "":
		return &messaging.Notification{
			Title: "",
			Body:  "",
		}
	default:
		return &messaging.Notification{
			Title: "Duitku Notification",
			Body:  "Duitku Notification",
		}
	}
}
