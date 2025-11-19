package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/health", HealthCheck)

	// --- User Registry ---
	app.Get("/user/:address", GetUser)
	app.Post("/user/verify", VerifyUserTx)

	// --- Reputation ---
	app.Get("/reputation/:address", GetReputation)
	app.Post("/reputation/add", AddReputationTx)
	app.Post("/reputation/sub", SubReputationTx)

	// --- Mentorship ---
	app.Post("/mentorship/request", RequestSessionTx)
	app.Post("/mentorship/accept", AcceptSessionTx)
	app.Post("/mentorship/complete", CompleteSessionTx)
	app.Post("/mentorship/feedback", GiveFeedbackTx)

	// --- Content Registry (IPFS-backed) ---
	app.Get("/content/:id", GetContent)
	app.Get("/content/public/list", ListPublicContent)
	app.Get("/content/uploads/:address", GetUserUploads)

	app.Post("/content/upload", UploadContent)

	// --- Content Access (ContentAccess.sol) ---
	app.Get("/content/access/:contentId/:viewer", CheckAccess)
	app.Post("/content/access/grant", GrantAccess)
	app.Post("/content/access/revoke", RevokeAccess)
}
