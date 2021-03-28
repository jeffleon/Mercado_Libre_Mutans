package main


import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jeffleon/mutansApi/database"
	"github.com/jeffleon/mutansApi/router"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

func main(){
	app := fiber.New()
	app.Use(cors.New())
	database.InitDatabase()
	defer database.DBConn.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		log.Printf("Defaulting to port %s", port)
	}

	router.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}