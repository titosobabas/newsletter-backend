package usecases

import (
	"fmt"
	"newsletter-backend/models"
	"newsletter-backend/types"
)

type NewsletterUseCase struct {
	model models.NewsletterModel
}

func NewNewsletterUseCase(model models.NewsletterModel) *NewsletterUseCase {
	return &NewsletterUseCase{model: model}
}

func (n *NewsletterUseCase) GetNewsletterEmails() ([]types.Newsletter, error) {
	newsletterEmails, err := n.model.GetEmails()
	if err != nil {
		return nil, fmt.Errorf("an error ocurred: [%v]", err)
	}
	return newsletterEmails, nil
}

func (n *NewsletterUseCase) SaveNewsletter(email string) (int, error) {
	_, err := n.model.EmailExists(email)
	if err != nil {
		return 3, fmt.Errorf("an error ocurred: [%s]", err)
	}
	code, _ := n.model.StoreData(email)
	return code, nil
}
