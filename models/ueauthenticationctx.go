package models

type UeAuthenticationCtx struct {
	AuthType           string
	Links              map[string]LinksValueSchema
	FivegAuthData      FivegAuthData
	ServingNetworkName string
}
