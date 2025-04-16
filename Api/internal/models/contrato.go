package models

type Contrato struct {
	Id           uint
	Nome         string `json:"nome"`
	CPF          string `json:"cpf"`
	Data_criacao string `json:"data_criacao"`
}
