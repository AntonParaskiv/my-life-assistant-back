package UserList

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
)

type List []*User.User

func New() (ul *List) {
	ul = &List{}
	return
}

func (list *List) AddUser(user *User.User) (err error) {
	list.Add(user)
	return
}

func (list *List) Add(user *User.User) *List {
	*list = append(*list, user)
	return list
}

func (list *List) GetUserByEmail(email string) (user *User.User) {
	key := list.searchUserByEmail(email)
	user = list.getUserByKey(key)
	return
}

func (list *List) IsUserExist(user *User.User) (isExist bool) {
	key := list.searchUserByEmail(user.Email())
	isExist = list.isKeyValid(key)
	return
}

func (list *List) searchUserByEmail(email string) (key int) {
	// set default not exist
	key = -1

	// search
	for listKey, user := range *list {
		if user.Email() == email {
			key = listKey
			break
		}
	}
	return
}

func (list *List) getUserByKey(key int) (user *User.User) {
	if !list.isKeyValid(key) {
		return
	}

	user = (*list)[key]
	return
}

func (list *List) isKeyValid(key int) (isValid bool) {
	if key < 0 {
		return
	}
	if key > list.lastKey() {
		return
	}
	isValid = true
	return
}

func (list *List) lastKey() (key int) {
	key = list.Len() - 1
	return
}

func (list *List) Len() (length int) {
	length = len(*list)
	return
}
