package handler

import "inventory-indra/repositories"

type TransactionHandler struct {
	repo *repositories.TransactionRepository
}

func NewTransactionHandler(repo *repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}