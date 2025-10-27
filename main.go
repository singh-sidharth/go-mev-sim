package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/singh-sidharth/go-mev-sim/blockbuilder"
	"github.com/singh-sidharth/go-mev-sim/mempool"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start mempool simulation
	mp := mempool.NewMempool()
	go mp.Run(ctx)

	// Start block builder
	bb := blockbuilder.NewBuilder(mp)
	go bb.Run(ctx)

	// HTTP API
	http.HandleFunc("/mempool", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mp.JSON()))
	})

	http.HandleFunc("/block", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(bb.JSON()))
	})

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
