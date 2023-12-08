package router

import (
	"fmt"
	"github.com/Sun-FreePort/Cities-and-Citizen/config"
	"github.com/Sun-FreePort/Cities-and-Citizen/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/storage/redis/v3"
	"github.com/gofiber/swagger"
	"os"
	"runtime"
	"time"
)

type Router struct {
	H *handler.Handler
}

// RegisterF2E 静态资源注册
func (r Router) RegisterF2E(app *fiber.App) {
	app.Static("/", "./public")

	app.Get("/swagger/*", swagger.HandlerDefault) // default
}

// RegisterB2E 无状态接口注册
func (r Router) RegisterB2E(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Ops! nothing happen...")
	})
}

// RegisterB2EAuth 鉴权接口注册
func (r Router) RegisterB2EAuth(app *fiber.App) {
	// 鉴权中间件 fixme 抽出去！抽出去！
	authKeyMiddleware := keyauth.New(keyauth.Config{
		AuthScheme: "Bearer",
		Validator: func(c *fiber.Ctx, key string) (bool, error) {
			if config.GetConfig("")["ENV"] == "test" && key == "test_token" {
				return true, nil
			}

			keyInfo := config.GetRedis().Get(fmt.Sprintf("auth:%s", "key"))
			if keyInfo == "" {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			config.GetRedis().Expire("key", 48*time.Hour)
			return true, nil
		},
	})

	// 路由
	square := app.Group("/square", authKeyMiddleware)
	square.Get("/info", r.H.SquareInfo)
	square.Post("/publish", r.H.SquarePublish)
}

// NewApp 新建一个应用
func (r Router) NewApp(configDict map[string]string, logFile *os.File) *fiber.App {
	isProd := configDict["ENV"] == "prod" || configDict["ENV"] == "production"

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
		Max:        45,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		Storage: storage,
	}))

	route := Router{
		H: handler.NewHandler(),
	}
	route.RegisterF2E(app)

	// 日志
	if logFile != nil {
		app.Use(logger.New(logger.Config{
			Format:     "${time} [${ip}:${port}] ${status} - ${method} ${path} ${resBody}\n",
			TimeFormat: time.DateTime,
			Output:     logFile,
		}))
	}

	route.RegisterB2E(app)

	route.RegisterB2EAuth(app)

	return app
}
