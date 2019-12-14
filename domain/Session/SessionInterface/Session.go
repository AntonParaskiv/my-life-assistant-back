package SessionInterface

import "github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"

type Session interface {
	Id() (id string)
	User() (user UserInterface.User)
	SetId(id string) Session
	SetUser(user UserInterface.User) Session
}
