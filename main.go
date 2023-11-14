package main

import (
	"fmt"
	"github.com/Sun-FreePort/Cities-and-Citizen/config"
	_ "github.com/Sun-FreePort/Cities-and-Citizen/docs"
	"github.com/Sun-FreePort/Cities-and-Citizen/handler"
	"github.com/Sun-FreePort/Cities-and-Citizen/router"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"time"
)

// @title Cities-and-Citizen API
// @version 1.0
// @description Wish nice day for you! 👋
// @termsOfService https://www.uiosun.com/policy/privacy
// @contact.name C&C API Support
// @contact.email uiosun@outlook.com
// @license.name GPL-3.0
// @license.url https://github.com/Sun-FreePort/Cities-and-Citizen/blob/main/LICENSE
// @host localhost:22042
// @BasePath /
func main() {
	configDict := config.GetConfig("")
	route := &router.Router{
		H: handler.NewHandler(),
	}
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", time.Now().Format(time.DateOnly)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app := route.NewApp(configDict, file)

	err = app.Listen(fmt.Sprintf(":%s", configDict["HTTP_PORT"]))
	if err != nil {
		panic(err)
	}
}
