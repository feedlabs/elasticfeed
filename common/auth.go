package common

import (
	auth "github.com/abbot/go-http-auth"
	"github.com/astaxie/beego/context"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/common/config"
)

var (
	a *auth.DigestAuth
)

func Auth(ctx *context.Context) *resource.Admin {
	if config.GetAuthType() == "basic" {
		return AuthBasic(ctx)
	} else if config.GetAuthType() == "digest" {
		return AuthDigest(ctx)
	}
	return nil
}

func SecretBasic(user, realm string) string {
	admin, err := resource.FindAdminByUsername(user)
	if err == nil {
		token := admin.Data
		return GetCrypt(token)
	}
	return ""
}

func SecretDigest(user, realm string) string {
	admin, err := resource.FindAdminByUsername(user)
	if err == nil {
		token := admin.Data
		return GetMd5(user + ":" + realm + ":" + token)
	}
	return ""
}

func AuthBasic(ctx *context.Context) *resource.Admin {
	authenticator := auth.NewBasicAuthenticator(config.GetAuthRealm(), SecretBasic)

	username := authenticator.CheckAuth(ctx.Request)
	if username == "" {
		authenticator.RequireAuth(ctx.ResponseWriter, ctx.Request)
	}

	admin, err := resource.FindAdminByUsername(username)
	if err == nil {
		return nil
	}

	return admin
}

func AuthDigest(ctx *context.Context) (admin *resource.Admin) {
	if a == nil {
		a = auth.NewDigestAuthenticator(config.GetAuthRealm(), SecretDigest)
	}

	username, authinfo := a.CheckAuth(ctx.Request)
	if username == "" {
		a.RequireAuth(ctx.ResponseWriter, ctx.Request)
	} else {

		var err error
		admin, err = resource.FindAdminByUsername(username)

		if err != nil {
			return admin
		}

		if !admin.IsWhitelisted(GetIP(ctx.Request)) {
			a.RequireAuth(ctx.ResponseWriter, ctx.Request)
		}

		ar := &auth.AuthenticatedRequest{Request: *ctx.Request, Username: username}
		if authinfo != nil {
			ctx.ResponseWriter.Header().Set("Authentication-Info", *authinfo)
		}

		if ar.Username == "" {
			return nil
		}
	}

	return admin
}

