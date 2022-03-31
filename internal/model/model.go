package model

import (
	"time"
)

type User struct {
	Username     string    `json:"username" gorm:"primary_key"`
	CreatedDate  time.Time `json:"createdDate"`
	LastLoggedIn time.Time `json:"lastLoggedIn"`
}

type Message struct {
	ID            string    `json:"id" gorm:"primary_key"`
	Sender        string    `json:"sender"`
	RecipientType string    `json:"recipientType"` // group or user
	RecipientName string    `json:"recipientName"`
	Body          string    `json:"body"`
	Subject       string    `json:"subject"`
	CreatedDate   time.Time `json:"createdDate"`
}

type Group struct {
	Name        string    `json:"name" gorm:"primary_key"`
	Creator     string    `json:"groupCreator,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
}

type GroupUser struct {
	Name     string `json:"groupname"`
	Username string `json:"usernames"`
}

type Reply struct {
	FromMessageID string `json:"fromMessageID"`
	ToMessageID   string `json:"toMessageID"`
}
