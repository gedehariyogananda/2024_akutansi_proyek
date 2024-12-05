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
const (
	SchemeMaterialStock   = "material-stock"
	SchemeSellableProduct = "sellable-product"
	SellableStock         = "sellable-stock"
)

func (i PushNotificationScheme) DefinePushNotificationMessages() *messaging.Notification {
	switch i {
	case SchemeMaterialStock:
		return &messaging.Notification{
			Title: "Stok Material Menipis",
			Body:  "",
		}
	case SchemeSellableProduct:
		return &messaging.Notification{
			Title: "Produk yang bisa dijual mau habis",
			Body:  "",
		}
	case SellableStock:
		return &messaging.Notification{
			Title: "Stok barang hampir habis",
			Body:  "",
		}
	default:
		return &messaging.Notification{
			Title: "Duitku Notification",
			Body:  "Duitku Notification",
		}
	}
}
