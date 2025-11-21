package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"backend_go/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// --- exported contract instances (initialized by LoadContracts) ---
var UserRegistry *bindings.UserRegistry
var Reputation *bindings.Reputation
var Mentorship *bindings.Mentorship
var ContentRegistry *bindings.ContentRegistry
var ContentAccess *bindings.ContentAccess
var Client *ethclient.Client

var ErrOperatorKeyMissing = errors.New("OPERATOR_PRIVATE_KEY environment variable not set")
var frontendSyncURL = "http://localhost:3000/api/onchain/sync"

func NotifyFrontendOnChainVerified(addr, tx string, block int64) error {
	payload := map[string]any{
		"address":     addr,
		"txHash":      tx,
		"blockNumber": block,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", frontendSyncURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	_, err := http.DefaultClient.Do(req)
	return err
}

// new helper: notify frontend about reputation changes
func NotifyFrontendReputationChanged(addr string, newScore int64, txHash string) error {
	payload := map[string]any{
		"event":    "reputation_changed",
		"address":  addr,
		"newScore": newScore,
		"txHash":   txHash,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", frontendSyncURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	_, err := http.DefaultClient.Do(req)
	return err
}

// Connect dials the Ethereum client.
func Connect(clientUrl string) {
	var err error
	Client, err = ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	if opKey := os.Getenv("OPERATOR_PRIVATE_KEY"); opKey != "" {
		pk := strings.TrimPrefix(opKey, "0x")
		priv, err := crypto.HexToECDSA(pk)
		if err == nil {
			addr := crypto.PubkeyToAddress(priv.PublicKey)
			bal, err := Client.BalanceAt(context.Background(), addr, nil)
			if err == nil {
				log.Println("Operator address:", addr.Hex(), "balance (wei):", bal.String())
			} else {
				log.Println("Failed to get operator balance:", err)
			}
		} else {
			log.Println("Invalid OPERATOR_PRIVATE_KEY:", err)
		}
	} else {
		log.Println("OPERATOR_PRIVATE_KEY not set")
	}

	log.Println("Connected RPC:", clientUrl)

	log.Println("Connected to Ethereum client")
}

// LoadContracts binds contract addresses to generated bindings.
func LoadContracts(urAddr, repAddr, mentAddr, crAddr, caAddr string) {
	log.Println("Using Reputation contract at:", repAddr)
	chainID, _ := Client.NetworkID(context.Background())
	log.Println("Connected RPC chain id:", chainID.String())

	var err error
	UserRegistry, err = bindings.NewUserRegistry(common.HexToAddress(urAddr), Client)
	if err != nil {
		log.Fatal(err)
	}
	Reputation, err = bindings.NewReputation(common.HexToAddress(repAddr), Client)
	if err != nil {
		log.Fatal(err)
	}
	Mentorship, err = bindings.NewMentorship(common.HexToAddress(mentAddr), Client)
	if err != nil {
		log.Fatal(err)
	}
	ContentRegistry, err = bindings.NewContentRegistry(common.HexToAddress(crAddr), Client)
	if err != nil {
		log.Fatal(err)
	}
	ContentAccess, err = bindings.NewContentAccess(common.HexToAddress(caAddr), Client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All contracts loaded successfully")
}

// --- watch helpers (optional) ---
func watchOpts() *bind.WatchOpts {
	start := uint64(0)
	return &bind.WatchOpts{
		Start:   &start,
		Context: context.Background(),
	}
}

func WatchMentorshipEvents() {
	ch := make(chan *bindings.MentorshipRequested)
	sub, err := Mentorship.WatchRequested(watchOpts(), ch, []*big.Int{}, []common.Address{}, []common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case evt := <-ch:
				log.Printf("New Mentorship Request: ID=%d, Student=%s, Mentor=%s", evt.Id, evt.Student.Hex(), evt.Mentor.Hex())
			case err := <-sub.Err():
				if err == nil {
					log.Println("Watcher closed")
					return
				}
				log.Println("Watcher error:", err)
				return
			}
		}
	}()
}

func WatchContentEvents() {
	ch := make(chan *bindings.ContentRegistryContentUploaded)
	sub, err := ContentRegistry.WatchContentUploaded(watchOpts(), ch, []common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case evt := <-ch:
				log.Printf("New Content Uploaded: ID=%d, Uploader=%s, Title=%s, Public=%v", evt.ContentId, evt.Uploader.Hex(), evt.Title, evt.IsPublic)
			case err := <-sub.Err():
				if err == nil {
					log.Println("Watcher closed")
					return
				}
				log.Println("Content watcher error:", err)
				return
			}
		}
	}()
}

func WatchUserEvents() {
	ch := make(chan *bindings.UserRegistryUserVerified)
	sub, err := UserRegistry.WatchUserVerified(watchOpts(), ch, []common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case evt := <-ch:
				log.Printf("User Verified: %s, Role=%d, Name=%s", evt.User.Hex(), evt.Role, evt.Name)
			case err := <-sub.Err():
				if err == nil {
					log.Println("Watcher closed")
					return
				}
				log.Println("User watcher error:", err)
				return
			}
		}
	}()
}

func WatchReputationEvents() {
	ch := make(chan *bindings.ReputationReputationChanged)
	sub, err := Reputation.WatchReputationChanged(watchOpts(), ch, []common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case evt := <-ch:
				log.Printf("Reputation Changed: Mentor=%s, Delta=%d, NewScore=%d", evt.Mentor.Hex(), evt.Delta.Int64(), evt.NewScore.Int64())
				// notify frontend/backend so persisted user records can be refreshed
				if err := NotifyFrontendReputationChanged(evt.Mentor.Hex(), evt.NewScore.Int64(), evt.Raw.TxHash.Hex()); err != nil {
					log.Println("Failed to notify frontend about reputation change:", err)
				}
			case err := <-sub.Err():
				if err == nil {
					log.Println("Watcher closed")
					return
				}
				log.Println("Reputation watcher error:", err)
				return
			}
		}
	}()
}

// --- low-level auth helper ---
func makeAuthFromPrivateKey(pkHex string) (*bind.TransactOpts, error) {
	priv, err := crypto.HexToECDSA(strings.TrimPrefix(pkHex, "0x"))
	if err != nil {
		return nil, err
	}

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Println("Failed to fetch network/chain id from RPC:", err)
		chainID = big.NewInt(11155111)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	if err != nil {
		return nil, err
	}
	log.Println("Using chain ID:", chainID.String())
	return auth, nil
}

// --- Read-only (call) helpers used by handlers ---

// GetUserInfo returns the user struct fields from the UserRegistry binding call
func GetUserInfo(addr common.Address) (map[string]interface{}, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}
	u, err := UserRegistry.GetUser(callOpts, addr)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"verified":    u.Verified,
		"role":        u.Role,
		"name":        u.Name,
		"rollNumber":  u.RollNumber,
		"program":     u.Program,
		"major":       u.Major,
		"pictureHash": u.PictureHash,
	}, nil
}

// GetContent fetches content by id (big.Int)
func GetContent(id *big.Int) (map[string]interface{}, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}
	c, err := ContentRegistry.GetContent(callOpts, id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"contentId": id.Int64(),
		"uploader":  c.Uploader.Hex(),
		"cid":       c.Cid,
		"title":     c.Title,
		"isPublic":  c.IsPublic,
		"ipfsUrl":   cidToHttp(c.Cid),
	}, nil
}

// ListPublicContent returns latest public contents up to a default limit (100)
func ListPublicContent() ([]map[string]interface{}, error) {
	const defaultLimit = 100
	total, err := GetTotalContents()
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return []map[string]interface{}{}, nil
	}
	out := []map[string]interface{}{}
	for i := total - 1; i >= 0 && len(out) < defaultLimit; i-- {
		item, err := GetContent(big.NewInt(i))
		if err != nil {
			log.Println("GetContent error:", err)
			continue
		}
		if isPub, ok := item["isPublic"].(bool); ok && isPub {
			out = append(out, item)
		}
	}
	return out, nil
}

// GetTotalContents returns total contents as an int64
func GetTotalContents() (int64, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}
	total, err := ContentRegistry.GetTotalContents(callOpts)
	if err != nil {
		return 0, err
	}
	return total.Int64(), nil
}

// GetUserUploads returns user's uploaded content items
func GetUserUploads(user common.Address) ([]map[string]interface{}, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}
	ids, err := ContentRegistry.GetUserUploads(callOpts, user)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]interface{}, 0, len(ids))
	for _, bi := range ids {
		item, err := GetContent(bi)
		if err != nil {
			log.Println("GetContent error (GetUserUploads):", err)
			continue
		}
		out = append(out, item)
	}
	return out, nil
}

// CheckAccess checks if viewer can access content (wrapper over ContentAccess.CanAccess)
func CheckAccess(contentId *big.Int, viewer common.Address) (bool, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}
	ok, err := ContentAccess.CanAccess(callOpts, contentId, viewer)
	if err != nil {
		return false, err
	}
	return ok, nil
}

// --- State-changing (tx) helpers used by handlers ---

// VerifyUserTx registers/verifies a user in UserRegistry
func VerifyUserTx(data *struct {
	PrivateKey string
	Address    string
	Role       uint8
	Name       string
	RollNo     string
	Program    string
	Major      string
	Pic        string
}) (string, error) {
	auth, err := makeAuthFromPrivateKey(data.PrivateKey)
	if err != nil {
		return "", err
	}
	tx, err := UserRegistry.VerifyUser(auth, common.HexToAddress(data.Address), data.Role, data.Name, data.RollNo, data.Program, data.Major, data.Pic)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// VerifyUserOnChain uses a server-side operator private key (from env)
// to call UserRegistry.VerifyUser for the given user address and metadata.
// Expects OPERATOR_PRIVATE_KEY in env as hex (with or without 0x).
func VerifyUserOnChain(data *struct {
	Address string
	Role    uint8
	Name    string
	RollNo  string
	Program string
	Major   string
	Pic     string
}) (string, error) {
	// read operator key from env
	opKey := os.Getenv("OPERATOR_PRIVATE_KEY")
	if opKey == "" {
		return "", ErrOperatorKeyMissing
	}

	// strip 0x if present and create auth
	privHex := strings.TrimPrefix(opKey, "0x")
	auth, err := makeAuthFromPrivateKey(privHex)
	if err != nil {
		return "", err
	}

	// call contract
	addr := common.HexToAddress(data.Address)
	tx, err := UserRegistry.VerifyUser(auth, addr, data.Role, data.Name, data.RollNo, data.Program, data.Major, data.Pic)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// UploadContentTx uploads content metadata to ContentRegistry
func UploadContentTx(privateKeyHex, cid, title string, isPublic bool) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := ContentRegistry.UploadContent(auth, cid, title, isPublic)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// GrantAccessTx grants viewer access to a content via ContentAccess contract
func GrantAccessTx(privateKeyHex string, contentId *big.Int, viewer common.Address) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := ContentAccess.Grant(auth, contentId, viewer)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// RevokeAccessTx revokes viewer access
func RevokeAccessTx(privateKeyHex string, contentId *big.Int, viewer common.Address) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := ContentAccess.Revoke(auth, contentId, viewer)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// Backwards-compatible names (older functions that existed in your repo)
func GrantContentAccessTx(privateKeyHex string, contentId int64, viewerHex string) (string, error) {
	return GrantAccessTx(privateKeyHex, big.NewInt(contentId), common.HexToAddress(viewerHex))
}

func RevokeContentAccessTx(privateKeyHex string, contentId int64, viewerHex string) (string, error) {
	return RevokeAccessTx(privateKeyHex, big.NewInt(contentId), common.HexToAddress(viewerHex))
}

func CanAccessContent(contentId int64, viewerHex string) (bool, error) {
	return CheckAccess(big.NewInt(contentId), common.HexToAddress(viewerHex))
}

// --- Mentorship tx helpers (already used by your handlers) ---

func RequestMentorshipTx(privateKeyHex string, mentorHex string) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := Mentorship.Request(auth, common.HexToAddress(mentorHex))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func AcceptSessionTx(privateKeyHex string, sessionId int64) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := Mentorship.Accept(auth, big.NewInt(sessionId))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func CompleteSessionTx(privateKeyHex string, sessionId int64) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := Mentorship.Complete(auth, big.NewInt(sessionId))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func GiveFeedbackTx(privateKeyHex string, sessionId int64, upvote bool) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	tx, err := Mentorship.GiveFeedback(auth, big.NewInt(sessionId), upvote)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// --- Reputation helpers (read / write wrappers) ---

func GetReputation(user common.Address) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}
	return Reputation.ScoreOf(callOpts, user)
}

func AddReputation(privateKeyHex string, mentor common.Address, amount *big.Int) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	log.Println("Sending tx from:", auth.From.Hex())

	// Read current operator so we can restore later
	currOp, err := Reputation.Operator(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return "", err
	}
	log.Println("Current operator:", currOp.Hex())

	// If auth.From is already the operator we can directly call Add
	needRestore := true
	if strings.EqualFold(currOp.Hex(), auth.From.Hex()) {
		needRestore = false
	}

	// 1) if needed, set operator to this auth (owner must be calling this)
	if needRestore {
		setOpTx, err := Reputation.SetOperator(auth, auth.From)
		if err != nil {
			return "", err
		}
		log.Println("Sent Reputation.SetOperator ->", setOpTx.Hash().Hex())

		// wait for mining
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		receipt, err := bind.WaitMined(ctx, Client, setOpTx)
		if err != nil {
			log.Println("WaitMined error for setOperator:", err)
			return setOpTx.Hash().Hex(), nil
		}
		if receipt == nil || receipt.Status == 0 {
			log.Println("setOperator tx failed or reverted")
			return setOpTx.Hash().Hex(), nil
		}
	}

	// 2) Call Reputation.add (now auth.From == operator)
	tx, err := Reputation.Add(auth, mentor, amount)
	if err != nil {
		// attempt best-effort restore if we changed operator
		if needRestore {
			_, _ = Reputation.SetOperator(auth, currOp)
		}
		return "", err
	}
	log.Println("Sent Reputation.Add tx:", tx.Hash().Hex())

	// wait for mining
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, Client, tx)
	if err != nil {
		log.Println("WaitMined error (Add):", err)
		// best-effort restore operator
		if needRestore {
			_, _ = Reputation.SetOperator(auth, currOp)
		}
		return tx.Hash().Hex(), nil
	}
	if receipt == nil {
		log.Println("No receipt returned from WaitMined (Add)")
	} else {
		log.Printf("Reputation Add mined. tx=%s status=%d logs=%d\n", tx.Hash().Hex(), receipt.Status, len(receipt.Logs))
		if receipt.Status == 0 {
			log.Println("Reputation.Add reverted. Check contract/caller funds/params.")
		}
	}

	// 3) restore operator back to previous value (if we changed it)
	if needRestore {
		restoreTx, err := Reputation.SetOperator(auth, currOp)
		if err != nil {
			log.Println("Failed to send restore setOperator:", err)
			// still return tx hash of add — don't fail hard
			return tx.Hash().Hex(), nil
		}
		log.Println("Sent Reputation.SetOperator(restore) ->", restoreTx.Hash().Hex())

		ctx2, cancel2 := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel2()
		_, err = bind.WaitMined(ctx2, Client, restoreTx)
		if err != nil {
			log.Println("WaitMined error for restore setOperator:", err)
		}
	}

	// read updated score from chain (best-effort)
	newScoreBig, err := GetReputation(mentor)
	var newScoreInt int64 = 0
	if err != nil {
		log.Println("Failed to read new reputation score after Add:", err)
	} else if newScoreBig != nil {
		newScoreInt = newScoreBig.Int64()
	}

	// notify frontend/backend of change (best-effort)
	if err := NotifyFrontendReputationChanged(mentor.Hex(), newScoreInt, tx.Hash().Hex()); err != nil {
		log.Println("Notify frontend error (AddReputation):", err)
	}

	return tx.Hash().Hex(), nil
}

func SubReputation(privateKeyHex string, mentor common.Address, amount *big.Int) (string, error) {
	auth, err := makeAuthFromPrivateKey(privateKeyHex)
	if err != nil {
		return "", err
	}
	log.Println("Sending tx from:", auth.From.Hex())

	// Read current operator so we can restore later
	currOp, err := Reputation.Operator(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return "", err
	}
	log.Println("Current operator:", currOp.Hex())

	needRestore := true
	if strings.EqualFold(currOp.Hex(), auth.From.Hex()) {
		needRestore = false
	}

	// 1) set operator to this auth if needed
	if needRestore {
		setOpTx, err := Reputation.SetOperator(auth, auth.From)
		if err != nil {
			return "", err
		}
		log.Println("Sent Reputation.SetOperator ->", setOpTx.Hash().Hex())

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		receipt, err := bind.WaitMined(ctx, Client, setOpTx)
		if err != nil {
			log.Println("WaitMined error for setOperator:", err)
			return setOpTx.Hash().Hex(), nil
		}
		if receipt == nil || receipt.Status == 0 {
			log.Println("setOperator tx failed or reverted")
			return setOpTx.Hash().Hex(), nil
		}
	}

	// 2) Call Reputation.sub
	tx, err := Reputation.Sub(auth, mentor, amount)
	if err != nil {
		if needRestore {
			_, _ = Reputation.SetOperator(auth, currOp)
		}
		return "", err
	}
	log.Println("Sent Reputation.Sub tx:", tx.Hash().Hex())

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, Client, tx)
	if err != nil {
		log.Println("WaitMined error (Sub):", err)
		if needRestore {
			_, _ = Reputation.SetOperator(auth, currOp)
		}
		return tx.Hash().Hex(), nil
	}
	if receipt == nil {
		log.Println("No receipt returned from WaitMined (Sub)")
	} else {
		log.Printf("Reputation Sub mined. tx=%s status=%d logs=%d\n", tx.Hash().Hex(), receipt.Status, len(receipt.Logs))
		if receipt.Status == 0 {
			log.Println("Reputation.Sub reverted. Check contract/caller funds/params.")
		}
	}

	// 3) restore operator
	if needRestore {
		restoreTx, err := Reputation.SetOperator(auth, currOp)
		if err != nil {
			log.Println("Failed to send restore setOperator:", err)
			return tx.Hash().Hex(), nil
		}
		log.Println("Sent Reputation.SetOperator(restore) ->", restoreTx.Hash().Hex())

		ctx2, cancel2 := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel2()
		_, err = bind.WaitMined(ctx2, Client, restoreTx)
		if err != nil {
			log.Println("WaitMined error for restore setOperator:", err)
		}
	}

	// read updated score from chain (best-effort)
	newScoreBig, err := GetReputation(mentor)
	var newScoreInt int64 = 0
	if err != nil {
		log.Println("Failed to read new reputation score after Sub:", err)
	} else if newScoreBig != nil {
		newScoreInt = newScoreBig.Int64()
	}

	// notify frontend/backend of change (best-effort)
	if err := NotifyFrontendReputationChanged(mentor.Hex(), newScoreInt, tx.Hash().Hex()); err != nil {
		log.Println("Notify frontend error (SubReputation):", err)
	}

	return tx.Hash().Hex(), nil
}

// --- misc utilities ---

func cidToHttp(cid string) string {
	if cid == "" {
		return ""
	}
	const ipfsPrefix = "ipfs://"
	if strings.HasPrefix(cid, ipfsPrefix) {
		cid = cid[len(ipfsPrefix):]
	}
	return "https://ipfs.io/ipfs/" + cid
}
