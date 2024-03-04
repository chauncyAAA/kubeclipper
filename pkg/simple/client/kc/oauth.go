package kc

import (
	"context"
	"encoding/json"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/oauth"
)

const (
	loginPath = "/oauth/login"
)

func (cli *Client) Login(ctx context.Context, body LoginRequest) (*oauth.Token, error) {
	serverResp, err := cli.post(ctx, loginPath, nil, body, JSONContentTypeHeader)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return nil, err
	}
	token := oauth.Token{}
	err = json.NewDecoder(serverResp.body).Decode(&token)
	return &token, err
}
