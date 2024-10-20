// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "github.com/Sanagiig/fox-admin-core/api/internal/handler/api"
	authority "github.com/Sanagiig/fox-admin-core/api/internal/handler/authority"
	base "github.com/Sanagiig/fox-admin-core/api/internal/handler/base"
	captcha "github.com/Sanagiig/fox-admin-core/api/internal/handler/captcha"
	configuration "github.com/Sanagiig/fox-admin-core/api/internal/handler/configuration"
	department "github.com/Sanagiig/fox-admin-core/api/internal/handler/department"
	dictionary "github.com/Sanagiig/fox-admin-core/api/internal/handler/dictionary"
	dictionarydetail "github.com/Sanagiig/fox-admin-core/api/internal/handler/dictionarydetail"
	menu "github.com/Sanagiig/fox-admin-core/api/internal/handler/menu"
	oauthprovider "github.com/Sanagiig/fox-admin-core/api/internal/handler/oauthprovider"
	position "github.com/Sanagiig/fox-admin-core/api/internal/handler/position"
	public_user "github.com/Sanagiig/fox-admin-core/api/internal/handler/public_user"
	publicapi "github.com/Sanagiig/fox-admin-core/api/internal/handler/publicapi"
	role "github.com/Sanagiig/fox-admin-core/api/internal/handler/role"
	token "github.com/Sanagiig/fox-admin-core/api/internal/handler/token"
	user "github.com/Sanagiig/fox-admin-core/api/internal/handler/user"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/core/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/job_database",
				Handler: base.InitJobDatabaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/mcms_database",
				Handler: base.InitMcmsDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: public_user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login_by_email",
				Handler: public_user.LoginByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login_by_sms",
				Handler: public_user.LoginBySmsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: public_user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register_by_email",
				Handler: public_user.RegisterByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register_by_sms",
				Handler: public_user.RegisterBySmsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/reset_password_by_email",
				Handler: public_user.ResetPasswordByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/reset_password_by_sms",
				Handler: public_user.ResetPasswordBySmsHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/create",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.GetUserByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/change_password",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/perm",
					Handler: user.GetUserPermCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/profile",
					Handler: user.GetUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/profile",
					Handler: user.UpdateUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/refresh_token",
					Handler: user.RefreshTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/access_token",
					Handler: user.AccessTokenHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role/create",
					Handler: role.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/update",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/delete",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.GetRoleByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/department/create",
					Handler: department.CreateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/update",
					Handler: department.UpdateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/delete",
					Handler: department.DeleteDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/list",
					Handler: department.GetDepartmentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department",
					Handler: department.GetDepartmentByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/create_or_update",
					Handler: authority.CreateOrUpdateApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/role",
					Handler: authority.GetApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/create_or_update",
					Handler: authority.CreateOrUpdateMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/role",
					Handler: authority.GetMenuAuthorityHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu/create",
					Handler: menu.CreateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/update",
					Handler: menu.UpdateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/delete",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/list",
					Handler: menu.GetMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/role/list",
					Handler: menu.GetMenuListByRoleHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/create",
					Handler: api.CreateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/update",
					Handler: api.UpdateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/delete",
					Handler: api.DeleteApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/list",
					Handler: api.GetApiListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api",
					Handler: api.GetApiByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/captcha/email",
				Handler: captcha.GetEmailCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/captcha/sms",
				Handler: captcha.GetSmsCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/configuration/create",
					Handler: configuration.CreateConfigurationHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/configuration/update",
					Handler: configuration.UpdateConfigurationHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/configuration/delete",
					Handler: configuration.DeleteConfigurationHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/configuration/list",
					Handler: configuration.GetConfigurationListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/configuration",
					Handler: configuration.GetConfigurationByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/configuration/system/list",
				Handler: publicapi.GetPublicSystemConfigurationListHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/create",
					Handler: dictionary.CreateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/update",
					Handler: dictionary.UpdateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/delete",
					Handler: dictionary.DeleteDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/list",
					Handler: dictionary.GetDictionaryListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary",
					Handler: dictionary.GetDictionaryByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionarydetail/create",
					Handler: dictionarydetail.CreateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionarydetail/update",
					Handler: dictionarydetail.UpdateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionarydetail/delete",
					Handler: dictionarydetail.DeleteDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionarydetail/list",
					Handler: dictionarydetail.GetDictionaryDetailListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionarydetail",
					Handler: dictionarydetail.GetDictionaryDetailByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/dict/:name",
					Handler: dictionarydetail.GetDictionaryDetailByDictionaryNameHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: oauthprovider.OauthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oauth/login/callback",
				Handler: oauthprovider.OauthCallbackHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/oauthprovider/create",
					Handler: oauthprovider.CreateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauthprovider/update",
					Handler: oauthprovider.UpdateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauthprovider/delete",
					Handler: oauthprovider.DeleteOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauthprovider/list",
					Handler: oauthprovider.GetOauthProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauthprovider",
					Handler: oauthprovider.GetOauthProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/position/create",
					Handler: position.CreatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/update",
					Handler: position.UpdatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/delete",
					Handler: position.DeletePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/list",
					Handler: position.GetPositionListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position",
					Handler: position.GetPositionByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token/create",
					Handler: token.CreateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/update",
					Handler: token.UpdateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/delete",
					Handler: token.DeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/list",
					Handler: token.GetTokenListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token",
					Handler: token.GetTokenByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/logout",
					Handler: token.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
