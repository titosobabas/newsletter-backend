package main

// $env:GOOS = "linux"
// go env
// go build -o main
import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
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

/*func HandleRequest(ctx context.Context) (string, error) {
	initEnvironmentVariables()
	db := initDatabase()
	repository := repositories.NewMysqlRepository(&db)
	modelNewsletter := models.NewNewsletterModel(repository)
	newsletterUsecase := usecases.NewNewsletterUseCase(*modelNewsletter)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	newsletterHttp(*e, *newsletterUsecase)
	e.Logger.Fatal(e.Start(":8000"))
	return "No me importan los eventos solo levantar mi api", nil
}*/

func initEnvironmentVariables() {
	os.Setenv("PORT_SERVER", os.Getenv("PORT_SERVER"))
	os.Setenv("DATABASE_NET", os.Getenv("DATABASE_NET"))
	os.Setenv("DATABASE_HOST", os.Getenv("DATABASE_HOST"))
	os.Setenv("DATABASE_USERNAME", os.Getenv("DATABASE_USERNAME"))
	os.Setenv("DATABASE_PASSWORD", os.Getenv("DATABASE_PASSWORD"))
	os.Setenv("DATABASE_NAME", os.Getenv("DATABASE_NAME"))
}

func main() {
	fmt.Println("Hola mundo!")
	initEnvironmentVariables()
	db := initDatabase()
	repository := repositories.NewMysqlRepository(&db)
	modelNewsletter := models.NewNewsletterModel(repository)
	newsletterUsecase := usecases.NewNewsletterUseCase(*modelNewsletter)
	newsletterTransport := transport.NewNewsletterTransport(*newsletterUsecase)

	e := echo.New()
	e.POST("/newsletter", newsletterTransport.SuscribeNewsletter)

	lambdaAdapter := &LambdaAdapter{Echo: e}
	fmt.Println("Loggeando algo")
	lambda.Start(lambdaAdapter.Handler)
}

/*func newsletterHttp(e echo.Echo, useCase usecases.NewsletterUseCase) {
	newsletterTransport := transport.NewNewsletterTransport(useCase)
	e.POST("/newsletter", newsletterTransport.SuscribeNewsletter)
}*/
