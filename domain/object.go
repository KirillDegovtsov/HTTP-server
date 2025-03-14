package domain

type TaskObject struct {
	Status string `json:"status"`
	Task   string `json:"task"`
	Result string `json:"result"`
	Id     string `json:"task_id"`
}

type UserObject struct {
	Id       string `json:"user_id"`
	Login    string `json:"username"`
	Password string `json:"password"`
}

type SessionObject struct {
	User_id    string `json:"user_id"`
	Session_id string `json:"session_id"`
}
