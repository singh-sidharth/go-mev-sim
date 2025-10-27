package mempool

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Tx struct {
	ID        string    `json:"id"`
	Sender    string    `json:"sender"`
	Amount    float64   `json:"amount"`
	Profit    float64   `json:"profit"`
	Timestamp time.Time `json:"timestamp"`
}

type Mempool struct {
	mu  sync.Mutex
	txs []Tx
}

func NewMempool() *Mempool {
	return &Mempool{txs: []Tx{}}
}

// Run simulates incoming transactions
func (m *Mempool) Run(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.mu.Lock()
			m.txs = append(m.txs, Tx{
				ID:        fmt.Sprintf("%d", rand.Intn(100000)),
				Sender:    fmt.Sprintf("user%d", rand.Intn(100)),
				Amount:    rand.Float64() * 10,
				Profit:    rand.Float64() * 5,
				Timestamp: time.Now(),
			})
			// Keep mempool small
			if len(m.txs) > 50 {
				m.txs = m.txs[len(m.txs)-50:]
			}
			m.mu.Unlock()
		}
	}
}

func (m *Mempool) GetAll() []Tx {
	m.mu.Lock()
	defer m.mu.Unlock()
	copied := make([]Tx, len(m.txs))
	copy(copied, m.txs)
	return copied
}

func (m *Mempool) JSON() string {
	data, _ := json.Marshal(m.GetAll())
	return string(data)
}
