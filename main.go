package main

import (
	"database/sql"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"newsletter-backend/database"
	"newsletter-backend/models"
	"newsletter-backend/repositories"
	"newsletter-backend/transport"
	"newsletter-backend/usecases"
)

func initDatabase() sql.DB {
	db, _ := database.Connect()
	return *db
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	db := initDatabase()
	repository := repositories.NewMysqlRepository(&db)
	modelNewsletter := models.NewNewsletterModel(repository)
	newsletterUsecase := usecases.NewNewsletterUseCase(*modelNewsletter)
	newsletterTransport := transport.NewNewsletterTransport(*newsletterUsecase)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.POST("/", newsletterTransport.SuscribeNewsletter)
	e.GET("/", newsletterTransport.GetListNewsletterEmails)

	lambdaAdapter := &LambdaAdapter{Echo: e}
	lambda.Start(lambdaAdapter.Handler)

}
