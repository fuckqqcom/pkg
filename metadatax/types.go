package metadatax

type UserInfo struct {
	Id       int64 `json:"id"`
	TenantId int64 `json:"tenant_id"`
	AppId    int64 `json:"app_id"`
	//OrganizationId int64 `json:"organization_id"`
	IsSuperuser bool `json:"is_superuser"`
}
