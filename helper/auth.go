package helper

import (
	auth "github.com/abbot/go-http-auth"
	"github.com/astaxie/beego/context"
	"github.com/feedlabs/api/resource"
)

var (
	a *auth.DigestAuth
)

func Auth(ctx *context.Context) *resource.Admin {
	if GetAuthType() == "basic" {
		return AuthBasic(ctx)
	} else if GetAuthType() == "digest" {
		return AuthDigest(ctx)
	}
	return nil
}

func SecretBasic(user, realm string) string {
	admin := GetAdminByName(user)
	if admin.Data != "" {
		token := admin.Data
		return GetCrypt(token)
	}
	return ""
}

func SecretDigest(user, realm string) string {
	if user == GetApiSuperuser() {
		token := GetApiSecret()
		return GetMd5(user + ":" + realm + ":" + token)
	}

	admin := GetAdminByName(user)
	if admin.Data != "" {
		token := admin.Data
		return GetMd5(user + ":" + realm + ":" + token)
	}
	return ""
}

func AuthBasic(ctx *context.Context) *resource.Admin {
	authenticator := auth.NewBasicAuthenticator(GetAuthRealm(), SecretBasic)

	username := authenticator.CheckAuth(ctx.Request)
	if username == "" {
		authenticator.RequireAuth(ctx.ResponseWriter, ctx.Request)
	}

	return GetAdminByName(username)
}

func AuthDigest(ctx *context.Context) (admin *resource.Admin) {
	if a == nil {
		a = auth.NewDigestAuthenticator(GetAuthRealm(), SecretDigest)
	}

	username, authinfo := a.CheckAuth(ctx.Request)
	if username == "" {
		a.RequireAuth(ctx.ResponseWriter, ctx.Request)
	} else {

		admin = GetAdminByName(username)

		if (IsSuperUser(admin) && GetIP(ctx.Request) != GetApiWhitelist()) ||
				!resource.Contains(GetAdminWhitelist(admin), GetIP(ctx.Request)) {

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

