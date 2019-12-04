package UserInterface

type User interface {
	Email() (email string)
	Password() (password string)
}
