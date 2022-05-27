package usecases

import (
	"fmt"
	"newsletter-backend/models"
)

type NewsletterUseCase struct {
	model models.NewsletterModel
}

func NewNewsletterUseCase(model models.NewsletterModel) *NewsletterUseCase {
	return &NewsletterUseCase{model: model}
}

/*func (n *NewsletterUseCase) EmailExistsAlready(email string) (int, error) {
	code, _ := n.model.StoreData(email)
	return code, nil
}*/

func (n *NewsletterUseCase) SaveNewsletter(email string) (int, error) {
	_, err := n.model.EmailExists(email)
	if err != nil {
		return 3, fmt.Errorf("an error ocurred: [%s]", err)
	}
	code, _ := n.model.StoreData(email)
	return code, nil
}
