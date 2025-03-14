package ram_storage

import (
	"my_project/domain"
	"my_project/repository"
	"my_project/usecases/service"
	"time"
)

type Object struct {
	task_data    map[string]domain.TaskObject
	user_data    map[string]domain.UserObject
	session_data map[string]domain.SessionObject
}

func NewObject() *Object {
	return &Object{
		task_data:    make(map[string]domain.TaskObject),
		user_data:    make(map[string]domain.UserObject),
		session_data: make(map[string]domain.SessionObject),
	}
}

func (rs *Object) GetStatus(task domain.TaskObject) (*string, error) {
	value, exists := rs.task_data[task.Id]
	if !exists {
		return nil, repository.NotFound
	}
	return &value.Status, nil
}

func (rs *Object) GetResult(task domain.TaskObject) (*string, error) {
	value, exists := rs.task_data[task.Id]
	if !exists {
		return nil, repository.NotFound
	}
	return &value.Result, nil
}

func (rs *Object) PostTask(task domain.TaskObject) error {
	if _, exists := rs.task_data[task.Id]; exists {
		return repository.AlreadyExists
	}
	rs.task_data[task.Id] = domain.TaskObject{
		Status: "in_progress",
		Task:   task.Task,
		Result: "",
		Id:     task.Id,
	}
	return nil
}

func (rs *Object) PutStatus(task domain.TaskObject) {
	time.Sleep(60 * time.Second)
	rs.task_data[task.Id] = domain.TaskObject{
		Status: "ready",
		Task:   task.Task,
		Result: "abobus",
		Id:     task.Id,
	}
}

func (rs *Object) PostRegister(user domain.UserObject) error {
	if _, exists := rs.user_data[user.Login]; exists {
		return repository.AlreadyExists
	}
	hash := service.MakeHashPassowrd(user.Password)
	if hash == "" {
		return repository.InternalError
	}
	rs.user_data[user.Login] = domain.UserObject{
		Id:       user.Id,
		Login:    user.Login,
		Password: hash,
	}
	return nil
}

func (rs *Object) PostLogin(user domain.UserObject, session domain.SessionObject) error {
	value, exists := rs.user_data[user.Login]
	if !exists {
		return repository.NotFound
	}
	if err := service.CheckValidPassword(value.Password, user.Password); err != nil {
		return repository.InvalidPassword
	}
	rs.session_data[session.Session_id] = domain.SessionObject{
		User_id:    value.Id,
		Session_id: session.Session_id,
	}
	return nil
}

func (rs *Object) DeleteSessionId(session domain.SessionObject) {
	time.Sleep(120 * time.Second)
	delete(rs.session_data, session.Session_id)
}

func (rs *Object) GetSession(session domain.SessionObject) bool {
	_, value := rs.session_data[session.Session_id]
	return value
}
