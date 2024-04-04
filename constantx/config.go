package constantx

const (
	Id             = "id"
	TenantId       = "tenant_id"
	AppId          = "app_id"
	UserId         = "user_id"
	AccountId      = "account_id"
	ArticleId      = "article_id"
	MaterialId     = "material_id"
	OrganizationId = "organization_id"
	MediaId        = "media_id"
	FactorId       = "factor_id"
	FactorConfigId = "factor_config_id"
	ProviderId     = "provider_id"
	AccessId       = "access_id"
	PlatformId     = "platform_id"
	AppKey         = "app_key"
	AppSecret      = "app_secret"

	OpenStatus      = "open_status"
	DeletedStatus   = "deleted_status"
	DraftedStatus   = "drafted_status"
	PublishedStatus = "published_status"
	CreatedTime     = "created_time"
	Code            = "code"
	Phone           = "phone"
	Level           = "level"
	Title           = "title"
	Content         = "content"
	// MediaImage 媒体文件:图片
	MediaImage = "image"
	// MediaVoice 媒体文件:声音
	MediaVoice = "voice"
	// MediaVideo 媒体文件:视频
	MediaVideo = "video"
	// MediaThumb 媒体文件:缩略图
	MediaThumb = "thumb"

	BaseAccess               = "base"
	BaseFactorPhone          = "phone"    //手机号码
	BaseFactorConfigPassword = "password" //密码登录
	BaseFactorConfigVerify   = "verify"   //验证码

	SocialAccess                          = "social"
	SocialFactorWechat                    = "wechat"
	SocialFactorConfigWechatOfficeAccount = "wechat_office_account" //公众号
	SocialFactorConfigWechatMini          = "wechat_mini"           //小程序
	SocialFactorConfigWechatPcQrcode      = "wechat_pc_qrcode"      // pc二维码
	EnterpriseAccess                      = "enterprise"
)
