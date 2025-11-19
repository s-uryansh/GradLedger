package main

import (
	"backend_go/api/handlers"
	"backend_go/services"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// connect to sepolia
	services.Connect("wss://sepolia.infura.io/ws/v3/c4bff9d0b6964e4c85bf89e71c0d4a53")

	services.LoadContracts(
		"0x623E59402bE01B511e373Bb68f67547BfD01b59e", // UserRegistry
		"0x8D284763B058A1536751A010AD63d11eafc949C2", // Reputation
		"0x5891F4255fE279f05fb1562B52E929254AdFC789", // Mentorship
		"0x371A83ec5190db8F93b4e868bca1E60090E11479", // ContentRegistry
		"0x8A50a537A11Da2981BE1041Dc4F3B7bfAe660dED", // ContentAccess
	)

	app := fiber.New()

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
