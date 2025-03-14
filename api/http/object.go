package http

import (
	"my_project/api/http/types"
	"my_project/domain"
	"my_project/usecases"
	"my_project/usecases/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Object struct {
	service usecases.Object
}

func NewHandler(service usecases.Object) *Object {
	return &Object{service: service}
}

func (s *Object) getHandlerResult(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if !s.service.GetSession(domain.SessionObject{Session_id: req.Auth_token}) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result, err := s.service.GetResult(domain.TaskObject{Id: req.Task_id})
	types.CreateObjectHandlerResponse(w, err, types.GetObjectHandlerResultResponse{Result: *result})
}

func (s *Object) getHandlerStatus(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if !s.service.GetSession(domain.SessionObject{Session_id: req.Auth_token}) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	status, err := s.service.GetStatus(domain.TaskObject{Id: req.Task_id})
	types.CreateObjectHandlerResponse(w, err, types.GetObjectHandlerStatusResponse{Status: *status})
}

func (s *Object) postHandlerTask(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if !s.service.GetSession(domain.SessionObject{Session_id: req.Auth_token}) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	uuid := service.MakeNewUuid()
	err = s.service.PostTask(domain.TaskObject{Id: uuid, Task: req.Task})
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	go s.service.PutStatus(domain.TaskObject{Id: uuid})
	types.CreateObjectHandlerResponse(w, nil, types.PostObjectHandlerTaskIdResponse{Task_id: uuid})
}

func (s *Object) postHandlerRegister(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostObjectHandlerUserRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	uuid := service.MakeNewUuid()
	err = s.service.PostRegister(domain.UserObject{Id: uuid, Login: req.Login, Password: req.Password})
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

func (s *Object) postHandlerLogin(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostObjectHandlerUserRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	session_id := service.MakeSessionId()
	if session_id == "" {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	err = s.service.PostLogin(domain.UserObject{Login: req.Login, Password: req.Password}, domain.SessionObject{Session_id: session_id})
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	go s.service.DeleteSessionId(domain.SessionObject{Session_id: session_id})
	types.CreateObjectHandlerResponse(w, nil, types.PostObjectHandlerAuthResponse{Auth_token: session_id})
}

func (s *Object) WithObjectHandlers(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
		r.Post("/", s.postHandlerTask)
	})

	r.Route("/result", func(r chi.Router) {
		r.Get("/", s.getHandlerResult)
	})

	r.Route("/status", func(r chi.Router) {
		r.Get("/", s.getHandlerStatus)
	})

	r.Route("/register", func(r chi.Router) {
		r.Post("/", s.postHandlerRegister)
	})

	r.Route("/login", func(r chi.Router) {
		r.Post("/", s.postHandlerLogin)
	})
}
