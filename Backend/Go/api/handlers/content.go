package handlers

import (
	"backend_go/services"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
)

func GetContent(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, ok := big.NewInt(0).SetString(idStr, 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	content, err := services.GetContent(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(content)
}

func ListPublicContent(c fiber.Ctx) error {
	list, err := services.ListPublicContent()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetUserUploads(c fiber.Ctx) error {
	addr := c.Params("address")
	if !common.IsHexAddress(addr) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid address"})
	}
	items, err := services.GetUserUploads(common.HexToAddress(addr))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func UploadContent(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		Cid        string `json:"cid"`
		Title      string `json:"title"`
		IsPublic   bool   `json:"isPublic"`
	})
	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	if in.PrivateKey == "" || in.Cid == "" || in.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing fields"})
	}
	tx, err := services.UploadContentTx(in.PrivateKey, in.Cid, in.Title, in.IsPublic)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"txHash": tx})
}

func CheckAccess(c fiber.Ctx) error {
	contentStr := c.Params("contentId")
	viewer := c.Params("viewer")
	if !common.IsHexAddress(viewer) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid viewer address"})
	}
	id, ok := big.NewInt(0).SetString(contentStr, 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid contentId"})
	}
	can, err := services.CheckAccess(id, common.HexToAddress(viewer))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"contentId": contentStr, "viewer": viewer, "canAccess": can})
}

func GrantAccess(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		ContentId  string `json:"contentId"`
		Viewer     string `json:"viewer"`
	})
	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	if !common.IsHexAddress(in.Viewer) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid viewer address"})
	}
	id, ok := big.NewInt(0).SetString(in.ContentId, 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid contentId"})
	}
	tx, err := services.GrantAccessTx(in.PrivateKey, id, common.HexToAddress(in.Viewer))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"txHash": tx})
}

func RevokeAccess(c fiber.Ctx) error {
	in := new(struct {
		PrivateKey string `json:"privateKey"`
		ContentId  string `json:"contentId"`
		Viewer     string `json:"viewer"`
	})
	if err := c.Bind().Body(in); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	if !common.IsHexAddress(in.Viewer) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid viewer address"})
	}
	id, ok := big.NewInt(0).SetString(in.ContentId, 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid contentId"})
	}
	tx, err := services.RevokeAccessTx(in.PrivateKey, id, common.HexToAddress(in.Viewer))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"txHash": tx})
}
