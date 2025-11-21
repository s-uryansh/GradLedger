package handlers

import (
	"backend_go/services"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	addr := c.Params("address")
	if addr == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing address"})
	}

	// Validate checksum format
	if !common.IsHexAddress(addr) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ethereum address"})
	}

	info, err := services.GetUserInfo(common.HexToAddress(addr))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(info)
}
func VerifyUserOnChainHandler(c fiber.Ctx) error {
	in := new(struct {
		Address string `json:"address"`
		Role    uint8  `json:"role"`
		Name    string `json:"name"`
		RollNo  string `json:"rollNo"`
		Program string `json:"program"`
		Major   string `json:"major"`
		Pic     string `json:"pic"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	if !common.IsHexAddress(in.Address) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid address"})
	}
	if in.Name == "" || in.RollNo == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name and rollNo required"})
	}
	if in.Role > 2 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid role"})
	}

	tx, err := services.VerifyUserOnChain(&struct {
		Address string
		Role    uint8
		Name    string
		RollNo  string
		Program string
		Major   string
		Pic     string
	}{
		Address: in.Address,
		Role:    in.Role,
		Name:    in.Name,
		RollNo:  in.RollNo,
		Program: in.Program,
		Major:   in.Major,
		Pic:     in.Pic,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"txHash": tx})
}

func VerifyUserTx(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		Address    string `json:"address"`
		Role       uint8  `json:"role"`
		Name       string `json:"name"`
		RollNo     string `json:"rollNo"`
		Program    string `json:"program"`
		Major      string `json:"major"`
		Pic        string `json:"pic"`
	})

	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	// Basic validation
	if in.PrivateKey == "" || in.Address == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing privateKey or address"})
	}
	if !common.IsHexAddress(in.Address) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ethereum address format"})
	}
	if in.Name == "" || in.RollNo == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name and rollNo required"})
	}
	if in.Role > 2 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid role"})
	}

	txHash, err := services.VerifyUserTx(&struct {
		PrivateKey string
		Address    string
		Role       uint8
		Name       string
		RollNo     string
		Program    string
		Major      string
		Pic        string
	}{
		PrivateKey: in.PrivateKey,
		Address:    in.Address,
		Role:       in.Role,
		Name:       in.Name,
		RollNo:     in.RollNo,
		Program:    in.Program,
		Major:      in.Major,
		Pic:        in.Pic,
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"txHash": txHash,
	})
}
