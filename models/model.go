package models

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

type BranchReq struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BranchRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
