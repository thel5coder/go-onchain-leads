# Go-Onchain-Leads 💎

A high-performance, real-time B2B Lead Generation engine built in **Golang**. This tool monitors the Ethereum blockchain to identify, validate, and unmask high-value developers the moment they launch new projects.

---

## 🚀 The Problem: Web3 Noise
The blockchain is a jungle of bots, scammers, and test transactions. For Web3 marketing agencies and audit firms, finding **real** humans with **real** budgets is like finding a needle in a haystack.

## 🎯 The Solution: Intelligent Filtering
`Go-Onchain-Leads` doesn't just scrape data; it validates **Skin in the Game**.
* **Contract Deployment Tracking:** Filters for the exact moment a new smart contract is initialized.
* **Economic Validation:** Automatically ignores "junk" deployments by filtering for transactions that spent significant Gas (default > 300,000 units).
* **Identity Unmasking:** Integrates with **ENS (Ethereum Name Service)** to resolve anonymous wallets into human-readable `.eth` domains.
* **Social Enrichment:** Extracts verified Twitter/X handles and Email addresses directly from on-chain text records.

---

## 🏗️ Technical Architecture
This project is built using **Clean Architecture (Hexagonal)** principles to ensure it is modular, testable, and scalable.

* **Domain Layer:** Pure business logic and data structures, independent of external libraries.
* **Usecase Layer:** Orchestrates the flow of data using **Interface Segregation** (Ports).
* **Adapters:** Specialized implementations for Ethereum (Geth), ENS Resolution, and CSV Storage.

### Project Structure
```text
├── cmd/                # Entry point (Dependency Injection)
├── internal/
│   ├── domain/         # Enterprise business objects (Leads, Identities)
│   ├── storage/        # File-system adapters (CSV Saver)
│   └── usecase/
│       └── leadscanner/# Core scanning & filtering logic (Ports & Services)```
```
## 🛠️ Getting Started
### Prerequisites
* Go 1.21+ 
* An Alchemy or Infura API Key (Ethereum Mainnet)

### Installation
**1. Clone the repository:** 
```text
git clone https://github.com/YOUR_USERNAME/go-onchain-leads.git
```
**2. Install dependencies:**
```text
go mod tidy
```
**3. Configure Connection:**

Update your RPC URL in cmd/main.go with your Alchemy API key.

**4. Run the engine:**
```text
go run cmd/main.go
```

## 📊 Output
Validated leads are automatically saved to premium_web3_leads.csv in real-time:

## 📜 License
MIT License - Feel free to use this for your own lead gen or as a foundation for Web3 tools.