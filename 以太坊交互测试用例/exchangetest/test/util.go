package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
	"context"
)

func WaitSendTx(ctx context.Context, b bind.DeployBackend, tx common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := b.TransactionReceipt(ctx, tx)
		if receipt != nil {
			return receipt, nil
		}
		if err != nil {
			fmt.Println("Receipt retrieval failed", "err", err)
		} else {
			fmt.Println("Transaction not yet mined")
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

