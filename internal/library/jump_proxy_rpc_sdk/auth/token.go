package auth

import (
	"context"
	"gin/internal/config"
)

type ClientToken struct {
	//
}

func (receive *ClientToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  config.Cfg.Rpc.AppId,
		"appKey": config.Cfg.Rpc.AppKey,
	}, nil
}

func (receive *ClientToken) RequireTransportSecurity() bool {
	return true
}
