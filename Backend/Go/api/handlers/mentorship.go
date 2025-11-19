package handlers

import (
	"backend_go/services"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
)

func RequestSessionTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		Mentor     string `json:"mentor"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if in.PrivateKey == "" || in.Mentor == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing privateKey or mentor"})
	}
	if !common.IsHexAddress(in.Mentor) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid mentor address"})
	}

	tx, err := services.RequestMentorshipTx(in.PrivateKey, in.Mentor)
	if err != nil {
		log.Println("RequestMentorshipTx error:", err)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}

func AcceptSessionTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		SessionId  int64  `json:"sessionId"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if in.PrivateKey == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing privateKey"})
	}
	if in.SessionId < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid sessionId"})
	}

	tx, err := services.AcceptSessionTx(in.PrivateKey, in.SessionId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}

func CompleteSessionTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		SessionId  int64  `json:"sessionId"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if in.PrivateKey == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing privateKey"})
	}
	if in.SessionId < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid sessionId"})
	}

	tx, err := services.CompleteSessionTx(in.PrivateKey, in.SessionId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}

func GiveFeedbackTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		SessionId  int64  `json:"sessionId"`
		Upvote     bool   `json:"upvote"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if in.PrivateKey == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing privateKey"})
	}
	if in.SessionId < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid sessionId"})
	}

	tx, err := services.GiveFeedbackTx(in.PrivateKey, in.SessionId, in.Upvote)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}
