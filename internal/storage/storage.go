package storage

import "time"

type User struct {
	FirstName string `json:"surname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Message struct {
	Authr   string `json:"author"`
	Content string `json:"content"`
	Created time.Time `json:"created"`
}

type Token struct {
	Token string `json:"token"`
}

type AuthStorage interface {
	Register(newUser *User) error
	Login(email string, password string) (string, error)
	GetUserInfo(email string) (*User, error)
}

type ChatStorage interface {
	SaveMessage(newMessage *Message) error 
	LoadChat() ([]Message,error)
}

