/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	pmempool "github.com/tylerztl/fabric-mempool/protos"
	"google.golang.org/grpc"
)

type MempoolClient interface {
	pmempool.MempoolClient
	Close() error
}

type mempoolClient struct {
	pmempool.MempoolClient
	conn *grpc.ClientConn
}

func NewMempoolClient(addr string) (MempoolClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &mempoolClient{
		pmempool.NewMempoolClient(conn),
		conn,
	}, nil

}

func (m *mempoolClient) Close() error {
	return m.conn.Close()
}
