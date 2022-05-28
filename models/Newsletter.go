package models

import "newsletter-backend/types"

type DatabaseRepository interface {
	GetEmails() ([]types.Newsletter, error)
	StoreData(email string) (int, error)
	EmailExists(email string) (int, error)
}

type NewsletterModel struct {
	repository DatabaseRepository
}

func NewNewsletterModel(repository DatabaseRepository) *NewsletterModel {
	return &NewsletterModel{repository: repository}
}

func (n *NewsletterModel) GetEmails() ([]types.Newsletter, error) {
	return n.repository.GetEmails()
}

func (n *NewsletterModel) EmailExists(email string) (int, error) {
	return n.repository.EmailExists(email)
}

func (n *NewsletterModel) StoreData(email string) (int, error) {
	return n.repository.StoreData(email)
}
