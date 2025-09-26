package services

import (
	"context"
	"log"
	"math/big"

	"backend_go/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var UserRegistry *bindings.UserRegistry
var Reputation *bindings.Reputation
var Mentorship *bindings.Mentorship
var ContentRegistry *bindings.ContentRegistry
var ContentAccess *bindings.ContentAccess
var Client *ethclient.Client

func Connect(clientUrl string) {
	var err error
	Client, err = ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	log.Println("Connected to Ethereum client")
}

func LoadContracts(urAddr, repAddr, mentAddr, crAddr, caAddr string) {
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
				log.Println("Mentorship watcher error:", err)
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
				log.Println("Content watcher error:", err)
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
				log.Println("User watcher error:", err)
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
			case err := <-sub.Err():
				log.Println("Reputation watcher error:", err)
			}
		}
	}()
}

func SubmitFeedback(auth *bind.TransactOpts, sessionID uint64, upvote bool) {
	go func() {
		tx, err := Mentorship.GiveFeedback(auth, big.NewInt(int64(sessionID)), upvote)
		if err != nil {
			log.Println("Feedback error:", err)
			return
		}
		log.Println("Feedback submitted, tx hash:", tx.Hash().Hex())
	}()
}
