package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func main() {
	db := initDatabase()
	repository := repositories.NewMysqlRepository(&db)
	// repository := repositories.NewPostgresqlRepository()
	modelNewsletter := models.NewNewsletterModel(repository)
	newsletterUsecase := usecases.NewNewsletterUseCase(*modelNewsletter)
	// response, err := newsletterUsecase.SaveNewsletter()
	// fmt.Println(response, err)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	newsletterHttp(*e, *newsletterUsecase)
	e.Logger.Fatal(e.Start(":1323"))

}

func newsletterHttp(e echo.Echo, useCase usecases.NewsletterUseCase) echo.Echo {
	newsletterTransport := transport.NewNewsletterTransport(useCase)
	e.POST("/newsletter", newsletterTransport.SuscribeNewsletter)
	return e
}
