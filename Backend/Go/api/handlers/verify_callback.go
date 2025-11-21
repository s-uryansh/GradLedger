package handlers

import (
	"backend_go/services"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
)

type VerifyCallbackIn struct {
	Address     string `json:"address"`
	TxHash      string `json:"txHash"`
	BlockNumber int64  `json:"blockNumber"`
}

func VerifyCallback(c fiber.Ctx) error {
	in := new(VerifyCallbackIn)

	// --- JSON validation ---
	if err := c.Bind().Body(in); err != nil {
		log.Println("[CALLBACK] Invalid JSON:", err)
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	log.Printf("[CALLBACK] Received: address=%s tx=%s block=%d",
		in.Address, in.TxHash, in.BlockNumber)

	if in.Address == "" || !common.IsHexAddress(in.Address) {
		log.Println("[CALLBACK] Invalid address")
		return c.Status(400).JSON(fiber.Map{"error": "invalid address"})
	}
	if in.TxHash == "" {
		log.Println("[CALLBACK] Missing txHash")
		return c.Status(400).JSON(fiber.Map{"error": "missing txHash"})
	}

	if services.Client != nil {
		receipt, err := services.Client.TransactionReceipt(context.Background(), common.HexToHash(in.TxHash))
		if err != nil {
			log.Println("[CALLBACK] Receipt not found yet:", err)
			return c.JSON(fiber.Map{"ok": true, "note": "tx pending"})
		}

		if receipt.Status != 1 {
			log.Printf("[CALLBACK] On-chain tx failed: %s", in.TxHash)
			return c.Status(400).JSON(fiber.Map{"error": "tx failed on chain"})
		}

		log.Printf("[CALLBACK] On-chain tx succeeded: %s", in.TxHash)
	}

	// --- Forward to your Next.js backend to update DB ---
	go func() {
		err := services.NotifyFrontendOnChainVerified(in.Address, in.TxHash, in.BlockNumber)
		if err != nil {
			log.Println("[CALLBACK] Frontend sync failed:", err)
		} else {
			log.Println("[CALLBACK] Frontend sync success")
		}
	}()

	return c.JSON(fiber.Map{"ok": true})
}
