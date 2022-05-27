package transport

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"newsletter-backend/types"
	"newsletter-backend/usecases"
)

type EmailBody struct {
	Email string `json:"email"`
}

type NewsletterTransport struct {
	usecase usecases.NewsletterUseCase
}

func NewNewsletterTransport(usecase usecases.NewsletterUseCase) *NewsletterTransport {
	return &NewsletterTransport{usecase: usecase}
}

func (nt *NewsletterTransport) SuscribeNewsletter(c echo.Context) error {

	response := types.HTTPResponse{Success: true, Response: ""}
	return c.JSON(http.StatusOK, response)
	/*newEmail := EmailBody{}
	if err := c.Bind(&newEmail); err != nil {
		response.Success = false
		response.Response = "Email is missing in the request"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Response = fmt.Sprintf("The email %s suscribed to the newsletter correctly!", newEmail.Email)
	_, err := nt.usecase.SaveNewsletter(newEmail.Email)
	if err != nil {
		response.Success = false
		response.Response = err.Error()
	}
	return c.JSON(http.StatusOK, response)*/
}
