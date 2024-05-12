package constantx

const (
	Id              = "id"
	TenantId        = "tenant_id"
	AppId           = "app_id"
	UserId          = "user_id"
	AccountId       = "account_id"
	ArticleId       = "article_id"
	MaterialId      = "material_id"
	OrganizationId  = "organization_id"
	MediaId         = "media_id"
	FactorId        = "factor_id"
	FactorKindId    = "factor_config_id"
	ProviderId      = "provider_id"
	AccessId        = "access_id"
	PlatformId      = "platform_id"
	ArticleKindId   = "article_kind_id"
	LanguageModelId = "language_model_id"
	Account         = "account"
	AppKey          = "app_key"
	AppSecret       = "app_secret"
	SocialAccountId = "social_account_id"
	IP              = "ip"
	UserInfo        = "userinfo"
	OpenStatus      = "open_status"
	DeletedStatus   = "deleted_status"
	DraftedStatus   = "drafted_status"
	PublishedStatus = "published_status"
	DefaultStatus   = "default_status"
	PrivateStatus   = "private_status"
	RetrievedStatus = "retrieved_status"
	ShowStatus      = "show_status"
	CreatedTime     = "created_time"
	Code            = "code"
	Name            = "name"
	Phone           = "phone"
	Email           = "email"
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

	BaseLoginMethod     = "base"
	BaseProviderPhone   = "phone"   //手机号码
	BaseProviderEmail   = "email"   //邮箱
	BaseProviderAccount = "account" //账号

	BaseIdentityProviderPassword = "password" //密码登录
	BaseIdentityProviderVerify   = "verify"   //验证码

	SocialLoginMethod                         = "social"
	SocialProviderWechat                      = "wechat"
	SocialIdentityProviderWechatOfficeAccount = "wechat_office_account" //公众号
	SocialIdentityProviderWechatMini          = "wechat_mini"           //小程序
	SocialIdentityProviderWechatPcQrcode      = "wechat_pc_qrcode"      // pc二维码
	EnterpriseLoginMethod                     = "enterprise"
)
