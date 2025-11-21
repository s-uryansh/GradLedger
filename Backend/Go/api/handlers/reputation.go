package handlers

import (
	"backend_go/services"
	"encoding/json"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
)

type repChangeIn struct {
	Mentor string `json:"mentor"`
	Amount int64  `json:"amount"`
}

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

func AddReputationTx(c fiber.Ctx) error {
	in := new(struct {
		Mentor string `json:"mentor"`
		Amount int64  `json:"amount"`
	})

	// parse JSON body using encoding/json (BodyParser not present in this build)
	if err := json.Unmarshal(c.Body(), in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	if !common.IsHexAddress(in.Mentor) || in.Amount <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	opKey := os.Getenv("OPERATOR_PRIVATE_KEY")
	if opKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "operator key missing"})
	}

	tx, err := services.AddReputation(
		opKey,
		common.HexToAddress(in.Mentor),
		big.NewInt(in.Amount),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}

func SubReputationTx(c fiber.Ctx) error {
	in := new(repChangeIn)
	if err := json.Unmarshal(c.Body(), in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if !common.IsHexAddress(in.Mentor) || in.Amount <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}
	opKey := os.Getenv("OPERATOR_PRIVATE_KEY")
	if opKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "operator key missing"})
	}
	tx, err := services.SubReputation(
		opKey,
		common.HexToAddress(in.Mentor),
		big.NewInt(in.Amount),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}
