package blockbuilder

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"github.com/singh-sidharth/go-mev-sim/mempool"
)

type Builder struct {
	mp       *mempool.Mempool
	topBlock []mempool.Tx
}

func NewBuilder(mp *mempool.Mempool) *Builder {
	return &Builder{mp: mp, topBlock: []mempool.Tx{}}
}

// Run periodically builds a candidate block with top N profitable transactions
func (b *Builder) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			txs := b.mp.GetAll()
			sort.Slice(txs, func(i, j int) bool {
				return txs[i].Profit > txs[j].Profit
			})
			top := 10
			if len(txs) < top {
				top = len(txs)
			}
			b.topBlock = txs[:top]
		}
	}
}

func (b *Builder) JSON() string {
	data, _ := json.Marshal(b.topBlock)
	return string(data)
}
