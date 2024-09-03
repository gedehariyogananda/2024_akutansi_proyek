package Models

import "time"

type StockType struct {
	ID        int
	TypeName  string
	CreatedAt time.Time
}
