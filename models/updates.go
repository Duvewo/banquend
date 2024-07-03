package models

import "encoding/json"

type UpdateModel struct {
	Type string // {"type": {"delete"|"update"|"new"}, "data": }
	Data json.RawMessage
}

type UpdateMessage struct{}
type UpdateTicket struct{}
