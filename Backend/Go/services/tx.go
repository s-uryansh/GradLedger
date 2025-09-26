package services

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestVerifyUser(privateKeyHex, userAddr string) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// ChainID must match your local node (Hardhat/Anvil usually = 31337)
	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	tx, err := UserRegistry.VerifyUser(
		auth,
		common.HexToAddress(userAddr),
		uint8(1), // Role.Student
		"Alice",
		"123",
		"CS",
		"Blockchain",
		"ipfs://hash",
	)
	if err != nil {
		log.Println("VerifyUser failed:", err)
		return
	}
	log.Println("VerifyUser tx:", tx.Hash().Hex())
}
