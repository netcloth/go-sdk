package rpc

import (
	"github.com/tendermint/tendermint/p2p"
	core_types "github.com/tendermint/tendermint/rpc/core/types"
)

type (
	ResultStatus struct {
		NodeInfo p2p.DefaultNodeInfo `json:"node_info"`
	}

	SyncStatus struct {
		SyncInfo core_types.SyncInfo `json:"sync_info"`
	}
)

func (c *client) GetStatus() (ResultStatus, error) {
	var (
		res ResultStatus
	)
	status, err := c.rpc.Status()
	if err != nil {
		return res, err
	} else {
		res.NodeInfo = status.NodeInfo
		return res, nil
	}
}

func (c *client) GetSyncStatus() (SyncStatus, error) {
	var (
		res SyncStatus
	)
	status, err := c.rpc.Status()
	if err != nil {
		return res, err
	} else {
		res.SyncInfo = status.SyncInfo
		return res, nil
	}
}
