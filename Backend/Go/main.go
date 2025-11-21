package main

import (
	"backend_go/api/handlers"
	"backend_go/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	// connect to sepolia
	services.Connect("wss://sepolia.infura.io/ws/v3/c4bff9d0b6964e4c85bf89e71c0d4a53")

	services.LoadContracts(
		"0x26FACdd7b5912537ddaAA0f0A19B3Ee0Eb4E0BBc", // UserRegistry
		"0x0d266b68CCDf64c34106f39f12B9Fec5a0D0ED86", // Reputation
		"0xC9B8139e244622E04411c614553899cC445Fd7B8", // Mentorship
		"0x0Df41985D12a3768FE67AE48C18190CE223BFf34", // ContentRegistry
		"0x9619f27Efed6Fa103478D4b17e6E0f8358A51C6d", // ContentAccess
	)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
	}))
	app.Use(func(c fiber.Ctx) error {
		println("[REQ]", c.Method(), c.Path())
		return c.Next()
	})

	handlers.RegisterRoutes(app)

	// event watchers
	go services.WatchUserEvents()
	go services.WatchContentEvents()
	go services.WatchMentorshipEvents()
	go services.WatchReputationEvents()

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
