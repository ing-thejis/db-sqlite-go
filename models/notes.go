package models

import(
	"time"
)

type Note struct{
	ID			int 		`json:"id, omitempty"`
	Title		string 		`json:"title"`
	Description	string 		`json:"description"`
	CreatedAt	time.Time 	`json:"create_at, omitempty"`
	UpdateAt	time.Time 	`json:"update_at, omitempty"`
}

