package UserInterface

type User interface {
	Email() (email string)
	Password() (password string)
	SetEmail(email string) User
	SetPassword(password string) User
}
