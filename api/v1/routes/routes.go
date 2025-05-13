package routes

import (
	"fmt"

	v1Handler "github.com/achmad-dev/simple-ewallet/api/v1/handler/v1"
	v1MiddleWare "github.com/achmad-dev/simple-ewallet/api/v1/middleware/v1"
	"github.com/achmad-dev/simple-ewallet/internal/pkg"
	"github.com/achmad-dev/simple-ewallet/internal/repository"
	"github.com/achmad-dev/simple-ewallet/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func ServerRouteV1(envpath string) {
	log := pkg.InitLog()
	log.Info("ServerRouteV1")

	cfg, err := pkg.NewConfig(envpath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	// Initialize the sqlx database
	postgreUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	// log
	log.Info("Connecting to database", postgreUrl)

	db, err := pkg.InitSqlDB(postgreUrl)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	// initialize bcrypt
	bcrypt := pkg.NewBcryptUtil(10)

	// initialize repository
	userRepo := repository.NewUserRepository(db)
	ewalletRepo := repository.NewEWalletRepository(db)
	// initialize service
	ewalletService := service.NewEWalletService(ewalletRepo, log)
	userService := service.NewUserService(userRepo, ewalletRepo, bcrypt, cfg.JwtSecret, log)

	// initialize handler
	authHandler := v1Handler.NewAuthHandler(userService)
	ewalletHandler := v1Handler.NewEWalletHandler(ewalletService)
	// initialize routes

	// initialize the fiber app
	app := fiber.New(
		fiber.Config{
			StrictRouting: true,
			Prefork:       true,
			AppName:       "Simple E-Wallet",
		},
	)
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
		}),
		fiberLog.New(
			fiberLog.Config{
				Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
				TimeFormat: "02-Jan-2006",
			},
		),
	)

	// initialize routes
	api := app.Group("/api/v1")

	// health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// auth routes
	auth := api.Group("/auth")
	auth.Post("/signup", authHandler.Signup)
	auth.Post("/login", authHandler.Login)
	// ewallet routes
	ewallet := api.Group("/ewallet")
	ewallet.Use(v1MiddleWare.AuthMiddleware(cfg.JwtSecret, userService))
	ewallet.Post("/add-balance", ewalletHandler.AddBalance)
	ewallet.Post("/withdraw", ewalletHandler.SubtractBalance)
	ewallet.Get("/get-wallet", ewalletHandler.GetWallet)

	// metrics
	// Metrics
	app.Get("/metrics", monitor.New(
		monitor.Config{
			Title: "Backend Metrics",
		},
	))

	log.Info("Starting server on port ", cfg.Port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))
	log.Info("Server stopped")
}
