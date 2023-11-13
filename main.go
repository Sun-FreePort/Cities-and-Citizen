package main

import (
	"fmt"
	"github.com/Sun-FreePort/Cities-and-Citizen/config"
	_ "github.com/Sun-FreePort/Cities-and-Citizen/docs"
	"github.com/Sun-FreePort/Cities-and-Citizen/handler"
	"github.com/Sun-FreePort/Cities-and-Citizen/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/storage/redis/v3"
	"os"
	"runtime"
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
	// 配置
	isProd := os.Getenv("ENV") == "prod" || os.Getenv("ENV") == "production"
	config.GetConfig("")

	// 应用
	app := fiber.New()
	app.Use(recover.New(recover.Config{
		EnableStackTrace: !isProd,
	}))

	// 头设置
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	// 限流
	storage := redis.New(redis.Config{
		Host:      "127.0.0.1",
		Port:      6379,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		Storage: storage,
	}))

	route := router.Router{
		H: handler.NewHandler(),
	}
	route.RegisterF2E(app)

	// 日志
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

	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}:${port}] ${status} - ${method} ${path} ${resBody}\n",
		TimeFormat: time.DateTime,
		Output:     file,
	}))

	route.RegisterB2E(app)

	handler.NewHandler()

	app.Listen(":22042")
}
