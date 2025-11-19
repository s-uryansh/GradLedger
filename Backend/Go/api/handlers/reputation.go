package handlers

import (
	"backend_go/services"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
)

// GET /reputation/:address
func GetReputation(c fiber.Ctx) error {
	addr := c.Params("address")
	if !common.IsHexAddress(addr) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid address"})
	}

	score, err := services.GetReputation(common.HexToAddress(addr))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"address": addr,
		"score":   score.String(),
	})
}

// POST /reputation/add
func AddReputationTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		Mentor     string `json:"mentor"`
		Amount     int64  `json:"amount"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if in.PrivateKey == "" || !common.IsHexAddress(in.Mentor) || in.Amount <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	tx, err := services.AddReputation(
		in.PrivateKey,
		common.HexToAddress(in.Mentor),
		big.NewInt(in.Amount),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}

// POST /reputation/sub
func SubReputationTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		Mentor     string `json:"mentor"`
		Amount     int64  `json:"amount"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if in.PrivateKey == "" || !common.IsHexAddress(in.Mentor) || in.Amount <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	tx, err := services.SubReputation(
		in.PrivateKey,
		common.HexToAddress(in.Mentor),
		big.NewInt(in.Amount),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}
