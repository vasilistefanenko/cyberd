package rpc

import "github.com/tendermint/tendermint/rpc/lib/types"

type ResultIndexStats struct {
	Height     uint64 `json:"height"`
	CidsCount  uint64 `json:"cidsCount"`
	LinksCount uint64 `json:"linksCount"`
	AccsCount  uint64 `json:"accsCount"`
}

func IndexStats(ctx *rpctypes.Context) (*ResultIndexStats, error) {
	stats := &ResultIndexStats{}
	stats.Height = uint64(cyberdApp.LastBlockHeight())
	stats.CidsCount = cyberdApp.CidsCount()
	stats.LinksCount = cyberdApp.LinksCount()
	stats.AccsCount = cyberdApp.AccsCount()
	return stats, nil
}
