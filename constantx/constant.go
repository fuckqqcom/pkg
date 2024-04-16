package constantx

const (
	Default        = -1
	Version        = 0
	IsNotDeleted   = 1  //未删除
	IsDeleted      = -1 //已删除
	IsNotDrafted   = -1 //未推送
	IsDrafted      = 1
	IsNotDefault   = 1
	IsDefault      = -1
	IsNotPublished = -1 //未发布
	IsPublished    = 1
	IsOpen         = 1 //开启
	IsNotOpen      = -1
	IsNotAudited   = -1
	IsAudited      = 1

	IsPrivate    = -1
	IsNotPrivate = 1

	IsNotRetrieved = -1
	IsRetrieved    = 1

	IsSuperuser    = 1
	IsNotSuperuser = -1
	IsStaff        = 1
	IsNotStaff     = 1

	AppSecretLen = 20
	AppKeyLen    = 12
	SaltLen      = 8
)
