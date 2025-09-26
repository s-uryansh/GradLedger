# Go Backend for GradLedger
This part of backend focuses on making project run transaction and requests in parallel using go routines.

Go is using bindings to connect to Smart contracts (testing was done on local node). These bindings were created using `abi` lib from go.

## Working:
- Each contract's binding is loaded in `services/blockchain.go` 
- Every binding has a watcher which watches request on that binding and run them using goroutine

### What is parallel:
All request (related to contracts) sent to projects are parallel for each account.

If transactions are on diff acc:

- Go backend runs these requests in parallel and sends them to the smart contracts concurrently. 
- The Ethereum node (or Hardhat) can mine them in any order, so this is genuine parallelism.

If transaction / request is on same account:
- Go backend can still launch them concurrently via goroutines.
- However, Ethereum enforces nonce order per account, so the transactions are mined sequentially.
- If a transaction with a lower nonce is pending, higher-nonce transactions will wait or fail with nonce too low.
# Conclusion
Parallelism is effective across different accounts.

Same-account transactions are concurrent at the Go backend level, but sequential at the blockchain level

---

# Run:
```bash
go mod tidy
go run main.go
```
