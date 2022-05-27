package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"newsletter-backend/database"
	"newsletter-backend/models"
	"newsletter-backend/repositories"
	"newsletter-backend/transport"
	"newsletter-backend/usecases"
	"os"
)

func initDatabase() sql.DB {
	db, _ := database.Connect()
	return *db
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func initEnvironmentVariables() {
	os.Setenv("PORT_SERVER", goDotEnvVariable("PORT_SERVER"))
	os.Setenv("DATABASE_NET", goDotEnvVariable("DATABASE_NET"))
	os.Setenv("DATABASE_HOST", goDotEnvVariable("DATABASE_HOST"))
	os.Setenv("DATABASE_USERNAME", goDotEnvVariable("DATABASE_USERNAME"))
	os.Setenv("DATABASE_PASSWORD", goDotEnvVariable("DATABASE_PASSWORD"))
	os.Setenv("DATABASE_NAME", goDotEnvVariable("DATABASE_NAME"))
}

func main() {
	initEnvironmentVariables()
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
