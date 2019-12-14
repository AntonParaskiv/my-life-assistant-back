package User

import "github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"

type User struct {
	email    string
	password string
}

func New() (u *User) {
	u = new(User)
	return u
}

func (u *User) SetEmail(email string) UserInterface.User {
	u.email = email
	return u
}

func (u *User) SetPassword(password string) UserInterface.User {
	u.password = password
	return u
}

func (u *User) Email() (email string) {
	email = u.email
	return
}

func (u *User) Password() (password string) {
	password = u.password
	return
}
