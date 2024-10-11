package svc

import (
	"github.com/Sanagiig/fox-admin-core/api/internal/config"
	i18n2 "github.com/Sanagiig/fox-admin-core/api/internal/i18n"
	"github.com/Sanagiig/fox-admin-core/api/internal/middleware"
	"github.com/Sanagiig/fox-admin-core/rpc/coreclient"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/captcha"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Trans     *i18n.Translator
	Redis     redis.UniversalClient
	Captcha   *base64Captcha.Captcha
	CoreRpc   coreclient.Core
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.CasbinDatabaseConf.Type, c.CasbinDatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Captcha:   captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		Trans:     trans,
		Redis:     rds,
		Casbin:    cbn,
		CoreRpc:   coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
	}
}
