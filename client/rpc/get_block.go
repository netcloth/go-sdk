package rpc

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *client) GetBlock(height *int64) (*ctypes.ResultBlock, error) {
	var (
		res *ctypes.ResultBlock
	)

	res, err := c.rpc.Block(height)
	if err != nil {
		return res, err
	} else {
		return res, nil
	}
}
