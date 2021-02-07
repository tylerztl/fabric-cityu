/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kafka

import (
	"context"
	"github.com/hyperledger/fabric/orderer/common/localconfig"
	"time"

	pb "github.com/hyperledger/fabric/protos/common"
	"google.golang.org/grpc"
)

func StartFetchTimer(mempool localconfig.MemPool) {
	conn, err := grpc.Dial(mempool.Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewMempoolClient(conn)

	logger.Infof("[mempool] start fetch transaction, addr=%v, block interval=%v", mempool.Address, mempool.Interval)
	c := time.Tick(mempool.Interval)
	for {
		select {
		case <-c:
			go func() {
				r, err := client.FetchTransactions(context.Background(), &pb.FetchTxsRequest{
					Sender: mempool.OrdererIdentity,
					TxNum:  10,
				})
				if err != nil {
					logger.Error("[mempool] failed to fetch transaction from mempool because = ", err)
					return
				}
				logger.Infof("[mempool] succeed to fetch transaction, txNum=%d, isEmpty=%v", r.TxNum, r.IsEmpty)
			}()
		}
	}
}
