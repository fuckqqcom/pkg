package metadatax

type UserInfo struct {
	Id       int64 `json:"id"`
	TenantId int64 `json:"tenant_id"`
	AppId    int64 `json:"app_id"`
}

type TokenInfo struct {
	Id   int64  `json:"id"`
	Data string `json:"data"`
}
