/*
Copyright 2017 Idealnaya rabota LLC
Licensed under Multy.io license.
See LICENSE for details
*/
package eth

import (
	"strings"

	pb "github.com/Multy-io/Multy-ETH-node-service/node-streamer"
	"github.com/Multy-io/Multy-back/store"
)

func (c *Client) txpoolTransaction(txHash string) {
	rawTx, err := c.Rpc.EthGetTransactionByHash(txHash)
	if err != nil {
		log.Errorf("c.Rpc.EthGetTransactionByHash:Get TX Err: %s", err.Error())
	}
	c.parseETHTransaction(*rawTx, -1, false)

	c.parseETHMultisig(*rawTx, -1, false)

	// add txpool record
	if rawTx.GasPrice.IsInt64() {
		c.AddToMempoolStream <- pb.MempoolRecord{
			Category: rawTx.GasPrice.Int64(),
			HashTX:   rawTx.Hash,
		}
	}

	if strings.ToLower(rawTx.To) == strings.ToLower(c.Multisig.FactoryAddress) {
		go func() {
			fi, err := parseFactoryInput(rawTx.Input)
			if err != nil {
				log.Errorf("txpoolTransaction:parseFactoryInput: %s", err.Error())
			}
			fi.TxOfCreation = txHash
			fi.FactoryAddress = c.Multisig.FactoryAddress
			fi.DeployStatus = int64(store.MultisigStatusDeployPending)
			c.NewMultisigStream <- *fi
		}()
	}

}
