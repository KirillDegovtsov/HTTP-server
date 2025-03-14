package service

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"my_project/domain"
	"my_project/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Object struct {
	repo repository.Object
}

func NewObject(repo repository.Object) *Object {
	return &Object{
		repo: repo,
	}
}

func (rs *Object) GetStatus(task domain.TaskObject) (*string, error) {
	return rs.repo.GetStatus(task)
}

func (rs *Object) GetResult(task domain.TaskObject) (*string, error) {
	return rs.repo.GetResult(task)
}

func (rs *Object) PostTask(task domain.TaskObject) error {
	return rs.repo.PostTask(task)
}

func (rs *Object) PutStatus(task domain.TaskObject) {
	rs.repo.PutStatus(task)
}

func MakeNewUuid() string {
	newuuid := uuid.New()
	return newuuid.String()
}

func (rs *Object) PostRegister(user domain.UserObject) error {
	return rs.repo.PostRegister(user)
}

func (rs *Object) PostLogin(user domain.UserObject, session domain.SessionObject) error {
	return rs.repo.PostLogin(user, session)
}

func (rs *Object) DeleteSessionId(session domain.SessionObject) {
	rs.repo.DeleteSessionId(session)
}

func (rs *Object) GetSession(session domain.SessionObject) bool {
	return rs.repo.GetSession(session)
}

func MakeSessionId() string {
	newsession := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, newsession); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(newsession)
}

func MakeHashPassowrd(password string) string {
	cost := 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func CheckValidPassword(oldPassword string, newPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword))
	return err
}
