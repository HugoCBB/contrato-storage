package models

type User struct {
	ID        uint
	Nome      string `json:"nome"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
	User_type string `json:"user_type"`
}
