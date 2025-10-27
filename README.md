# Go MEV Transaction Simulator

A lightweight simulator of a blockchain mempool and a candidate block builder implemented in Go, with a real-time frontend dashboard in JavaScript. Demonstrates low-latency transaction processing, concurrency in Go, and transaction prioritization by “profit” — core concepts relevant to DEX and Layer 1 blockchain infrastructure.

---

## Features

**Backend (Go)**  

- Simulates a mempool receiving random transactions with varying profit values.  
- Concurrent block builder selects top N profitable transactions every second.  
- Demonstrates Go concurrency using goroutines and mutexes.  
- Exposes HTTP API endpoints:  
  - `/mempool` → list all pending transactions  
  - `/block` → list top transactions for candidate block  

**Frontend (JavaScript Dashboard)**  

- Displays mempool and candidate block transactions in real-time.  
- Updates every second to simulate live blockchain activity.  
- Shows transaction ID, sender, amount, profit, and timestamp.  

---

## Motivation

Built to explore MEV-aware transaction ordering and production-ready Go microservices, this project mirrors the logic behind transaction prioritization and block building in DEXs, without needing Solidity or actual blockchain deployment.

---

## Project Structure

```text
go-mev-sim/
├── main.go         # Starts mempool, block builder, HTTP server
├── mempool/
│ └── mempool.go    # Handles transaction generation and storage
├── blockbuilder/
│ └── builder.go    # Selects top transactions by profit
├── frontend/
│ ├── index.html    # Simple dashboard interface
│ └── script.js     # Fetches data from backend and renders tables
├── go.mod
└── README.md
```

## Getting Started

### Run the Go backend

```bash
go run main.go
```

### Open the frontend dashboard

Open `frontend/index.html` in a browser. It fetches data from `http://localhost:8080` and updates live.

## Key Points

- Concurrency & Performance: Uses goroutines and mutexes for efficient, concurrent processing.

- MEV Awareness: Demonstrates prioritization of transactions by “profit,” analogous to real-world block-building strategies.

- Production-Oriented Design: Modular code structure, HTTP API, and frontend dashboard illustrate full-stack engineering practices.
