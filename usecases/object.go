package usecases

import (
	"my_project/domain"
)

type Object interface {
	GetStatus(task domain.TaskObject) (*string, error)
	GetResult(task domain.TaskObject) (*string, error)
	PostTask(task domain.TaskObject) error
	PutStatus(task domain.TaskObject)
	PostRegister(user domain.UserObject) error
	PostLogin(user domain.UserObject, session domain.SessionObject) error
	DeleteSessionId(session domain.SessionObject)
	GetSession(session domain.SessionObject) bool
}
