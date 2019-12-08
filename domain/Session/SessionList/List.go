package SessionList

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
)

type List []SessionInterface.Session

func New() (list *List) {
	list = &List{}
	return
}

func (list *List) AddSession(session SessionInterface.Session) (err error) {
	list.Add(session)
	return
}

func (list *List) Add(session SessionInterface.Session) *List {
	*list = append(*list, session)
	return list
}

func (list *List) GetSessionById(id string) (session SessionInterface.Session) {
	key := list.searchSessionById(id)
	session = list.getSessionByKey(key)
	return
}

func (list *List) IsSessionIdExist(session SessionInterface.Session) (isExist bool) {
	key := list.searchSessionById(session.Id())
	isExist = list.isKeyValid(key)
	return
}

func (list *List) searchSessionById(id string) (key int) {
	// set default not exist
	key = -1

	// search
	for listKey, session := range *list {
		if session.Id() == id {
			key = listKey
			break
		}
	}
	return
}

func (list *List) getSessionByKey(key int) (session SessionInterface.Session) {
	if !list.isKeyValid(key) {
		return
	}

	session = (*list)[key]
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
