package models

import "time"

type Contrato struct {
	Id           uint
	Nome         string    `json:"nome"`
	CPF          string    `json:"cpf"`
	Data_criacao time.Time `json:"data_criacao"`
}
