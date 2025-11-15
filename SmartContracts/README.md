# Blockchain Smart Contracts

Smart contracts for **Mentorship & Content Management Platform**. Manages users, content, access, mentorship sessions, and mentor reputation.  

---

## Contracts Overview

| Contract | Purpose | Key Features / Functions |
|----------|---------|--------------------------|
| **UserRegistry** | Manage verified users and roles | Roles: Student, Alumni, Faculty<br>`verifyUser(...)`, `isVerified(address)`, `getUser(address)` |
| **ContentRegistry** | Store and retrieve content | IPFS-based content (`cid`), public/private<br>`uploadContent(...)`, `getContent(id)`, `getUserUploads(user)`, `getTotalContents()` |
| **ContentAccess** | Access control for private content | Whitelist mechanism<br>`grant(contentId, viewer)`, `revoke(contentId, viewer)`, `canAccess(contentId, viewer)` |
| **Reputation** | Track mentor reputation | Only `Mentorship` can modify<br>`add(mentor, amt)`, `sub(mentor, amt)`, `scoreOf(user)` |
| **Mentorship** | Manage mentorship sessions | Request/accept/complete/cancel sessions<br>Feedback affects reputation<br>`request(mentor)`, `accept(id)`, `complete(id)`, `cancel(id)`, `giveFeedback(id, upvote)`, `getSession(id)`, `totalSessions()` |

---

## Deployment & Usage

### 1. Install dependencies
```bash
npm install
```

### 2. Compile contracts
```bash
npx hardhat compile
```

### 3. Run tests
```bash
npx hardhat test
```

### 4. Deploy locally
```bash
npx hardhat node
npx hardhat run scripts/deployAll.js --network localhost
```

**Suggested deployment order:**  
1. `UserRegistry` → verify users  
2. `ContentRegistry` → upload content  
3. `ContentAccess` → link to `ContentRegistry`  
4. `Reputation` → set operator to `Mentorship`  
5. `Mentorship` → link to `UserRegistry` & `Reputation`  

---

## Solidity version
- Solidity version: `^0.8.28`